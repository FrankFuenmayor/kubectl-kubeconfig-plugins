package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

const missingArgErrorMessage = `Missing namespace

Usage: 
	kubectl set-namespace <namespace>

`

type KConfig map[string]interface{}
type KContext map[interface{}]interface{}

func main() {

	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Print(missingArgErrorMessage)
		os.Exit(1)
	}

	selectedNamespace := flag.Arg(0)

	uhd, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	kubeConfigPath := path.Join(uhd, ".kube", "config")

	config, err := ioutil.ReadFile(kubeConfigPath)

	if err != nil {
		panic(err)
	}

	var out KConfig

	err = yaml.Unmarshal(config, &out)

	if err != nil {
		panic(err)
	}

	currentContext := out["current-context"]

	fmt.Printf(`Context "%v" default namespace is now "%v"\n`, currentContext, selectedNamespace)

	ctxs, ok := out["contexts"].([]interface{})

	if !ok {
		panic(errors.New("unexpected"))
	}

	var selectedContext KContext

	for _, ctx := range ctxs {

		ctxMap := ctx.(KContext)

		if currentContext == NameOf(ctxMap) {
			selectedContext = ctxMap["context"].(KContext)
		}
	}

	if selectedContext == nil {
		panic(errors.New(fmt.Sprintf("context %v not found", currentContext)))
	}

	selectedContext["namespace"] = selectedNamespace

	file, err := os.OpenFile(kubeConfigPath, os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	err = yaml.NewEncoder(writer).Encode(out)

	if err != nil {
		panic(err)
	}
}

func NameOf(contexts KContext) string {
	return contexts["name"].(string)
}
