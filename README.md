# code-generator
根据输入自动生成kubernetes CRD clientset、informers、listers、deepcopy文件。\
依赖：golang、code-generator 
## 使用
init:初始化项目，自动创建golang项目，并生成CRD clientset、informers、listers、deepcopy文件。\
add:在当前golang项目基础上,新增group或version的CRD clientset、informers、listers、deepcopy文件。
## 参数
```shell
Usage of api:
    --debug                       Code Generator Debug
-d, --domain string               Api Domain
-g, --group string                Api Group
-k, --kubernetes version string   Kubernetes Version
-m, --model string                Resource Model Name
-v, --version string              Api Version
```


## example
初始化项目example，添加group为one，版本为v1，domain为example.com，k8s版本为v0.20.0的CRD

```shell
code-generator init -m example -k v0.20.0 -g one -v v1 -d example.com
```

添加group为two，版本为v1，domain为example.com，k8s版本为v0.20.0的CRD

```shell
code-generator add -k v0.20.0 -g two -v v1 -d example.com
```

添加group为two，版本为v2，domain为example.com，k8s版本为v0.20.0的CRD

```shell
code-generator add -k v0.20.0 -g two -v v2 -d example.com
```

生成效果

```
.
├── apis
│   ├── one
│   │   └── v1
│   │       ├── doc.go
│   │       ├── types.go
│   │       └── zz_generated.deepcopy.go
│   └── two
│       ├── v1
│       │   ├── doc.go
│       │   ├── types.go
│       │   └── zz_generated.deepcopy.go
│       └── v2
│           ├── doc.go
│           ├── types.go
│           └── zz_generated.deepcopy.go
├── generated
│   ├── one
│   │   └── v1
│   │       ├── clientset
│   │       │   └── versioned
│   │       │       ├── clientset.go
│   │       │       ├── doc.go
│   │       │       ├── fake
│   │       │       │   ├── clientset_generated.go
│   │       │       │   ├── doc.go
│   │       │       │   └── register.go
│   │       │       ├── scheme
│   │       │       │   ├── doc.go
│   │       │       │   └── register.go
│   │       │       └── typed
│   │       │           └── one
│   │       │               └── v1
│   │       │                   ├── doc.go
│   │       │                   ├── fake
│   │       │                   │   ├── doc.go
│   │       │                   │   ├── fake_foo.go
│   │       │                   │   └── fake_one_client.go
│   │       │                   ├── foo.go
│   │       │                   ├── generated_expansion.go
│   │       │                   └── one_client.go
│   │       ├── informers
│   │       │   └── externalversions
│   │       │       ├── factory.go
│   │       │       ├── generic.go
│   │       │       ├── internalinterfaces
│   │       │       │   └── factory_interfaces.go
│   │       │       └── one
│   │       │           ├── interface.go
│   │       │           └── v1
│   │       │               ├── foo.go
│   │       │               └── interface.go
│   │       └── listers
│   │           └── one
│   │               └── v1
│   │                   ├── expansion_generated.go
│   │                   └── foo.go
│   └── two
│       ├── v1
│       │   ├── clientset
│       │   │   └── versioned
│       │   │       ├── clientset.go
│       │   │       ├── doc.go
│       │   │       ├── fake
│       │   │       │   ├── clientset_generated.go
│       │   │       │   ├── doc.go
│       │   │       │   └── register.go
│       │   │       ├── scheme
│       │   │       │   ├── doc.go
│       │   │       │   └── register.go
│       │   │       └── typed
│       │   │           └── two
│       │   │               └── v1
│       │   │                   ├── doc.go
│       │   │                   ├── fake
│       │   │                   │   ├── doc.go
│       │   │                   │   ├── fake_foo.go
│       │   │                   │   └── fake_two_client.go
│       │   │                   ├── foo.go
│       │   │                   ├── generated_expansion.go
│       │   │                   └── two_client.go
│       │   ├── informers
│       │   │   └── externalversions
│       │   │       ├── factory.go
│       │   │       ├── generic.go
│       │   │       ├── internalinterfaces
│       │   │       │   └── factory_interfaces.go
│       │   │       └── two
│       │   │           ├── interface.go
│       │   │           └── v1
│       │   │               ├── foo.go
│       │   │               └── interface.go
│       │   └── listers
│       │       └── two
│       │           └── v1
│       │               ├── expansion_generated.go
│       │               └── foo.go
│       └── v2
│           ├── clientset
│           │   └── versioned
│           │       ├── clientset.go
│           │       ├── doc.go
│           │       ├── fake
│           │       │   ├── clientset_generated.go
│           │       │   ├── doc.go
│           │       │   └── register.go
│           │       ├── scheme
│           │       │   ├── doc.go
│           │       │   └── register.go
│           │       └── typed
│           │           └── two
│           │               └── v2
│           │                   ├── doc.go
│           │                   ├── fake
│           │                   │   ├── doc.go
│           │                   │   ├── fake_foo.go
│           │                   │   └── fake_two_client.go
│           │                   ├── foo.go
│           │                   ├── generated_expansion.go
│           │                   └── two_client.go
│           ├── informers
│           │   └── externalversions
│           │       ├── factory.go
│           │       ├── generic.go
│           │       ├── internalinterfaces
│           │       │   └── factory_interfaces.go
│           │       └── two
│           │           ├── interface.go
│           │           └── v2
│           │               ├── foo.go
│           │               └── interface.go
│           └── listers
│               └── two
│                   └── v2
│                       ├── expansion_generated.go
│                       └── foo.go
├── go.mod
├── go.sum
└── hack
    ├── boilerplate.go.txt
    └── tools.go
```

