package yamlutils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Write(filepath string, obj interface{}) error {
	updatedData, err := yaml.Marshal(obj)
	err = ioutil.WriteFile(filepath, updatedData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read(filepath string, obj interface{}) error {

	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, obj)

	if err != nil {
		return err
	}

	return nil

}
