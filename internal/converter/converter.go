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
	data, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}

	// Определим входной формат по расширению
	var fromFormat string
	if strings.HasSuffix(inPath, ".json") {
		fromFormat = "json"
	} else if strings.HasSuffix(inPath, ".yaml") || strings.HasSuffix(inPath, ".yml") {
		fromFormat = "yaml"
	} else {
		return errors.New("unsupported input file format")
	}

	var intermediate interface{}

	switch fromFormat {
	case "json":
		if err := json.Unmarshal(data, &intermediate); err != nil {
			return err
		}
	case "yaml":
		if err := yaml.Unmarshal(data, &intermediate); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown input format: %s", fromFormat)
	}

	var out []byte
	switch toFormat {
	case "json":
		out, err = json.MarshalIndent(intermediate, "", "  ")
	case "yaml":
		out, err = yaml.Marshal(intermediate)
	default:
		return fmt.Errorf("unsupported output format: %s", toFormat)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(outPath, out, 0644)
}
