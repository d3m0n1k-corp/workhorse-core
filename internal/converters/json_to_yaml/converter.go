package json_to_yaml

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/types"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var mockableYamlMarshal = yaml.Marshal

type JsonToYamlConverter struct {
	config JsonToYamlConfig
}

func (j *JsonToYamlConverter) Apply(input any) (any, error) {
	logrus.Tracef("JsonToYamlConverter: Apply called with input %v", input)
	in_data, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid input type")
	}

	var data any
	err := json.Unmarshal([]byte(in_data), &data)
	if err != nil {
		return nil, err
	}
	logrus.Tracef("JsonToYamlConverter: Unmarshalled JSON data: %v", data)
	out, err := mockableYamlMarshal(data)
	if err != nil {
		logrus.Errorf("JsonToYamlConverter: Error marshalling to YAML: %v", err)
		return nil, err
	}
	logrus.Tracef("JsonToYamlConverter: Marshalled YAML data: %s", string(out))
	return string(out), nil
}

func (j *JsonToYamlConverter) InputType() string {
	return types.JSON
}

func (j *JsonToYamlConverter) OutputType() string {
	return types.YAML
}
