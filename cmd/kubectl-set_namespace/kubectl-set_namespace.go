package main

import (
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

const usage = `Usage: 

	kubectl set-namespace [namespace]`

var MissingNamespaceArgument = errors.New("Missing namespace argument")
var TooManyArguments = errors.New("Too many arguments")

func main() {

	configFlags := genericclioptions.NewConfigFlags(true)

	cmd := &cobra.Command{
		Use: "set-namespace",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return MissingNamespaceArgument
			}
			if len(args) > 1 {
				return TooManyArguments
			}

			config, err := configFlags.ToRawKubeConfigLoader().RawConfig()

			if err != nil {
				return err
			}

			currentContext := config.Contexts[config.CurrentContext]
			currentContext.Namespace = args[0]
			configPath := clientcmd.NewDefaultPathOptions()

			if err = clientcmd.ModifyConfig(configPath, config, true); err != nil {
				return err
			}

			color.Bold.Printf("✅  Context %v updated\n", config.CurrentContext)
			return nil
		},
	}

	configFlags.AddFlags(cmd.Flags())
	cmd.SetUsageTemplate(usage)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
