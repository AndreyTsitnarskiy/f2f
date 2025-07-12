package converter

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
)

// unmarshalTOML парсит TOML-данные в map[string]interface{}
func unmarshalTOML(data []byte) (interface{}, error) {
	var temp map[string]interface{}
	if _, err := toml.Decode(string(data), &temp); err != nil {
		return nil, fmt.Errorf("error decoding TOML: %v", err)
	}
	return temp, nil
}

// marshalTOML сериализует map[string]interface{} в TOML
func marshalTOML(data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("error encoding TOML: %v", err)
	}
	return buf.Bytes(), nil
}
