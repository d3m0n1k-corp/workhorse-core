package json_to_yaml

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/types"

	"gopkg.in/yaml.v3"
)

var mockableYamlMarshal = yaml.Marshal

type JsonToYamlConverter struct {
	config JsonToYamlConfig
}

func (j *JsonToYamlConverter) Apply(input any) (any, error) {
	in_data, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid input type")
	}

	var data any
	err := json.Unmarshal([]byte(in_data), &data)
	if err != nil {
		return nil, err
	}
	out, err := mockableYamlMarshal(data)
	if err != nil {
		return nil, err
	}
	return string(out), nil
}

func (j *JsonToYamlConverter) InputType() string {
	return types.JSON
}

func (j *JsonToYamlConverter) OutputType() string {
	return types.YAML
}
