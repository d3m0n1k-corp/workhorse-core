package yaml_to_json

import (
	"encoding/json"
	"fmt"
	"strings"
	"workhorse-core/internal/common/types"

	"gopkg.in/yaml.v3"
)

var mockableJsonMarshalIndent = json.MarshalIndent

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

	var indent string
	if y.config.IndentType == "space" {
		indent = strings.Repeat(" ", y.config.IndentSize)
	} else {
		indent = strings.Repeat("\t", y.config.IndentSize)
	}

	out, err := mockableJsonMarshalIndent(data, "", indent)
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
