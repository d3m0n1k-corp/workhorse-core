package yaml_to_json

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/types"

	"gopkg.in/yaml.v3"
)

var mockableJsonMarshal = json.Marshal

type YamlToJsonConverter struct {
	config YamlToJsonConfig
}

func (y *YamlToJsonConverter) Apply(input any) (any, error) {
	in_data, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid input type")
	}

	var data any
	err := yaml.Unmarshal([]byte(in_data), &data)
	if err != nil {
		return nil, err
	}

	out, err := mockableJsonMarshal(data)
	if err != nil {
		return nil, err
	}
	return string(out), nil
}

func (y *YamlToJsonConverter) InputType() string {
	return types.YAML
}

func (y *YamlToJsonConverter) OutputType() string {
	return types.JSON
}
