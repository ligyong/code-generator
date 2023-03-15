package option

import "github.com/spf13/pflag"

type CreateOptions struct {
	Model             string
	KubernetesVersion string
	Group             string
	ApiVersion        string
	Domain            string
	Debug             bool
}

func NewCreateOption() CreateOptions {
	return CreateOptions{}
}

func (options *CreateOptions) AddFlags(args *pflag.FlagSet) {
	args.StringVarP(&options.Model, "model", "m", "", "Resource Model Name")
	args.StringVarP(&options.KubernetesVersion, "kubernetes version", "k", "", "Kubernetes Version")
	args.StringVarP(&options.Group, "group", "g", "", "Api Group")
	args.StringVarP(&options.ApiVersion, "version", "v", "", "Api Version")
	args.StringVarP(&options.Domain, "domain", "d", "", "Api Domain")
	args.BoolVarP(&options.Debug, "debug", "", false, "Code Generator Debug")
}
