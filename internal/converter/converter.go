package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"strings"

	"gopkg.in/yaml.v3"
)

func ConvertFile(inPath, outPath, toFormat string) error {
	inData, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}

	// Определим входной формат по расширению
	var fromFormat string
	if strings.HasSuffix(inPath, ".json") {
		fromFormat = "json"
	} else if strings.HasSuffix(inPath, ".yaml") || strings.HasSuffix(inPath, ".yml") {
		fromFormat = "yaml"
	} else if strings.HasSuffix(inPath, ".toml") {
		fromFormat = "toml"
	} else {
		return errors.New("unsupported input file format")
	}

	var intermediateData interface{}

	switch fromFormat {
	case "json":
		err = json.Unmarshal(inData, &intermediateData)
	case "yaml":
		err = yaml.Unmarshal(inData, &intermediateData)
	case "toml":
		intermediateData, err = unmarshalTOML(inData)
	default:
		return errors.New("unsupported input format")
	}

	var out []byte
	switch toFormat {
	case "json":
		out, err = json.MarshalIndent(intermediateData, "", "  ")
	case "yaml":
		out, err = yaml.Marshal(intermediateData)
	case "toml":
		dataMap, ok := intermediateData.(map[string]interface{})
		if !ok {
			return errors.New("cannot encode TOML: unexpected data format")
		}
		out, err = marshalTOML(dataMap)
	default:
		return fmt.Errorf("unsupported output format: %s", toFormat)
	}

	if fromFormat == toFormat {
		return fmt.Errorf("source and target formats are the same (%s). No conversion needed", toFormat)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(outPath, out, 0644)
}
