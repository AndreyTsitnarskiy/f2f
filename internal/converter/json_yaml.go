package converter

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func JsonToYaml(inFile, outFile string) error {
	data, err := os.ReadFile(inFile)
	if err != nil {
		return err
	}

	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, yamlData, 0644)
}

func YamlToJson(inFile, outFile string) error {
	data, err := os.ReadFile(inFile)
	if err != nil {
		return err
	}

	var obj interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, jsonData, 0644)
}
