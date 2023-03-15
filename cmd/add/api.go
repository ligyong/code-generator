package add

import (
	"bufio"
	"code-generator/cmd/option"
	"code-generator/pkg/exec"
	"code-generator/pkg/template"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

func AddCommand() *cobra.Command {
	createOptions := option.NewCreateOption()
	optionSet := pflag.NewFlagSet("api", pflag.ContinueOnError)
	initCmd := &cobra.Command{
		Use:   "add resource api",
		Short: "add",
		Long:  "add resource api",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := optionSet.Parse(args); err != nil {
				return err
			}

			if err := run(createOptions); err != nil {
				return err
			}

			return nil
		},
		DisableFlagParsing: true,
	}

	createOptions.AddFlags(optionSet)

	return initCmd
}

// run
func run(options option.CreateOptions) error {
	//1、创建golang项目
	model, err := getModel()
	if err != nil {
		return err
	}

	options.Model = model

	downloadDependency(options.KubernetesVersion)

	if err := template.Generate(options.Group, options.Domain, options.ApiVersion); err != nil {
		return err
	}

	vendor()

	do(options)

	return nil
}

func getModel() (string, error) {
	file, err := os.OpenFile("go.mod", os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	if strings.Contains(string(line), "module") {
		args := strings.Split(string(line), " ")
		if len(args) < 2 {
			return "", errors.New("can't discriminate go.mod module")
		}

		return args[1], nil
	}

	return "", errors.New("can't discriminate go.mod module")

}

func downloadDependency(version string) error {
	exec.New(fmt.Sprintf("go get k8s.io/client-go@%s", version)).Run()
	exec.New(fmt.Sprintf("go get k8s.io/apimachinery@%s", version)).Run()
	exec.New(fmt.Sprintf("go get k8s.io/code-generator@%s", version)).Run()
	return nil
}

func vendor() {
	exec.New("go mod tidy").Run()
	exec.New("go mod vendor").Run()

	exec.New("chmod +x ./vendor/k8s.io/code-generator/generate-groups.sh").Run()
}

func do(options option.CreateOptions) {
	path, _ := os.Getwd()

	debug := ""
	if options.Debug {
		debug = "-v 10"
	}

	exec.New(fmt.Sprintf("./vendor/k8s.io/code-generator/generate-groups.sh all %s/generated/%s/%s %s/apis %s:%s --go-header-file %s/hack/boilerplate.go.txt --output-base %s/../ %s",
		options.Model,
		options.Group,
		options.ApiVersion,
		options.Model,
		options.Group,
		options.ApiVersion,
		path,
		path,
		debug,
	)).Run()
}
