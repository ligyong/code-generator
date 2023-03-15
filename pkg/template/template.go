package template

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const DOC_TEMPLATE = "" +
	"// +k8s:deepcopy-gen=package\n" +
	"// +groupName={{ .Groups }}.{{ .Domain }}\n\n" +
	"// {{ .Version }}版本的api包\n" +
	"package {{ .Version }}"

type DocTemplate struct {
	Groups  string `json:"groups"`
	Domain  string `json:"domain"`
	Version string `json:"version"`
}

func generateDoc(group, domain, version string) error {
	Doc := DocTemplate{
		Groups:  group,
		Domain:  domain,
		Version: version,
	}

	//解析魔板
	t, err := template.New("").Parse(DOC_TEMPLATE)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf(DOC_PATH, group, version))
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	if err = t.Execute(f, Doc); err != nil {
		return err
	}

	return nil
}

const TYPES_TEMPLATE = "" +
	"package {{.Version}}\r\n\r\n" +
	"import (\r\n\t" +
	"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"\r\n" +
	")\r\n\r\n" +
	"// +genclient\r\n" +
	"// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\r\n\r\n" +
	"// Foo is a specification for a Foo resource\r\n" +
	"type Foo struct {\r\n\t" +
	"metav1.TypeMeta   `json:\",inline\"`\r\n\t" +
	"metav1.ObjectMeta `json:\"metadata,omitempty\"`\r\n\r\n\t" +
	"Spec   FooSpec   `json:\"spec\"`\r\n\t" +
	"Status FooStatus `json:\"status\"`\r\n" +
	"}\r\n\r\n" +
	"// FooSpec is the spec for a Foo resource\r\n" +
	"type FooSpec struct {\r\n\t" +
	"DeploymentName string`json:\"deploymentName\"`\r\n\t" +
	"Replicas       *int32 `json:\"replicas\"`\r\n" +
	"}\r\n\r\n" +
	"// FooStatus is the status for a Foo resource\r\n" +
	"type FooStatus struct {\r\n\t" +
	"AvailableReplicas int32 `json:\"availableReplicas\"`\r\n" +
	"}\r\n\r\n" +
	"// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\r\n\r\n" +
	"// FooList is a list of Foo resources\r\n" +
	"type FooList struct {\r\n\t" +
	"metav1.TypeMeta `json:\",inline\"`\r\n\t" +
	"metav1.ListMeta `json:\"metadata\"`\r\n\r\n\t" +
	"Items []Foo `json:\"items\"`\r\n" +
	"}"

func generateTypes(group, version string) error {
	type Types struct {
		Version string
	}

	//解析魔板
	t, err := template.New("").Parse(TYPES_TEMPLATE)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf(TYPES_PATH, group, version))
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	if err = t.Execute(f, Types{Version: version}); err != nil {
		return err
	}

	return nil
}

const TOOLS_TEMPLATE = "" +
	"// +build tools\r\n\r\n" +
	"/*\r\n" +
	"Copyright 2019 The Kubernetes Authors.\r\n\r\n" +
	"Licensed under the Apache License, Version 2.0 (the \"License\");\r\n" +
	"you may not use this file except in compliance with the License.\r\n" +
	"You may obtain a copy of the License at\r\n\r\n" +
	"    http://www.apache.org/licenses/LICENSE-2.0\r\n\r\n" +
	"Unless required by applicable law or agreed to in writing, software\r\n" +
	"distributed under the License is distributed on an \"AS IS\" BASIS,\r\n" +
	"WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\r\n" +
	"See the License for the specific language governing permissions and\r\n" +
	"limitations under the License.\r\n" +
	"*/\r\n\r\n" +
	"// This package imports things required by build scripts, to force `go mod` to see them as dependencies\r\n" +
	"package tools\r\n\r\n" +
	"import _ \"k8s.io/code-generator\"" +
	""

const BOILERPLATE_TEMPLATE = `
/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
`

func createFile(path, body string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.WriteString(body)
	if err != nil {
		return err
	}

	f.Close()
	return nil
}

const (
	API_DIR          = "./apis/%s/%s"
	HACK_DIR         = "./hack"
	DOC_PATH         = "./apis/%s/%s/doc.go"   //	./apis/group/version/doc.go
	TYPES_PATH       = "./apis/%s/%s/types.go" //	./apis/group/version/types.go
	TOOLS_PATH       = "./hack/tools.go"
	BOILERPLATE_PATH = "./hack/boilerplate.go.txt"
)

func Generate(group, domain, version string) error {
	//创建文件夹
	//apis/group/version
	if err := os.MkdirAll(fmt.Sprintf(API_DIR, group, version), 0660); err != nil {
		fmt.Println("1", err)

		return err
	}
	//hack
	if err := os.MkdirAll(HACK_DIR, 0660); err != nil {
		fmt.Println("2", err)

		return err
	}
	//生成doc.go
	if err := generateDoc(group, domain, version); err != nil {
		fmt.Println("3", err)

		return err
	}
	//生成types.go
	if err := generateTypes(group, version); err != nil {
		fmt.Println("4", err)

		return err
	}
	//生成tools.go
	if err := createFile(TOOLS_PATH, TOOLS_TEMPLATE); err != nil {
		fmt.Println("5", err)

		return err
	}
	//生成boilerplate.go.txt
	if err := createFile(BOILERPLATE_PATH, BOILERPLATE_TEMPLATE); err != nil {
		fmt.Println("6", err)

		return err
	}

	return nil
}
