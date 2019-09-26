package aws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Eks struct {
	Args []string
}

func (e Eks) ListClusters() ([]string, error) {

	stdOut, errOut, err := execute("list-clusters", e.Args)

	if err != nil {
		if errOut != nil {
			fmt.Println(errOut.String())
		}
		return nil, err
	}

	var output struct {
		Clusters []string
	}

	if err := json.Unmarshal(stdOut.Bytes(), &output); err != nil {
		return nil, err
	}

	return output.Clusters, nil

}

func (e Eks) UpdateKubeconfig(clusters []string) error {
	for _, cluster := range clusters {
		updatedArgs := append([]string{"--name", cluster, "--alias", cluster}, e.Args...)
		_, stdErr, err := execute("update-kubeconfig", updatedArgs)
		if err != nil {
			fmt.Println(stdErr.String())
			return err
		}
	}
	return nil
}

func execute(subcommand string, args []string) (*bytes.Buffer, *bytes.Buffer, error) {

	cmdArgs := append([]string{"eks", subcommand}, args...)

	command := exec.Command("aws", cmdArgs...)

	stdBuffer := bytes.NewBufferString("")
	stdErrBuffer := bytes.NewBufferString("")

	command.Stdout = stdBuffer
	command.Stderr = stdErrBuffer

	if err := command.Run(); err != nil {
		return stdBuffer, stdErrBuffer, err
	}

	return stdBuffer, stdErrBuffer, nil
}
