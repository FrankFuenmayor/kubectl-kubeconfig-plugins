package main

import (
	"flag"
	"github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/pkg/kubeconfig"
	"github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/pkg/utils/yamlutils"
	"github.com/gookit/color"
	"os"
	"path"
)

const missingArgErrorMessage = `Missing namespace
Usage: 
	kubectl set-namespace [namespace]

`

func main() {

	validateArgs()

	selectedNamespace := flag.Arg(0)

	kubeConfigPath, err := resolveKubeConfigPath()

	if err != nil {
		panic(err)
	}

	kubeConfig, err := newKubeConfig(kubeConfigPath)

	if err != nil {
		panic(err)
	}

	if err := kubeConfig.UpdateCurrentContextNamespace(selectedNamespace); err != nil {
		panic(err)
	}

	if err := yamlutils.Write(kubeConfigPath, kubeConfig); err != nil {
		panic(err)
	}

	color.Bold.Printf("📃 %v updated\n", kubeConfigPath)
	color.Bold.Printf("✅ %v namespace: %v\n", kubeConfig.CurrentContextName(), selectedNamespace)
}

func validateArgs() {
	flag.Parse()
	if flag.NArg() == 0 {
		color.Bold.Print(missingArgErrorMessage)
		os.Exit(1)
	}
}

func resolveKubeConfigPath() (string, error) {

	if path := os.Getenv("KUBECONFIG"); path != "" {
		return path, nil
	}

	uhd, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path.Join(uhd, ".kube", "config"), nil

}

func newKubeConfig(configPath string) (kubeconfig.KubeConfig, error) {

	var out kubeconfig.KubeConfig

	if err := yamlutils.Read(configPath, &out); err != nil {
		return nil, err
	}

	return out, nil
}
