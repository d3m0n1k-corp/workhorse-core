package json_to_yaml

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var mockableYamlMarshal = yaml.Marshal

type JsonToYamlConverter struct {
	config JsonToYamlConfig
}

// Apply - string-based execution for single operations
func (j *JsonToYamlConverter) Apply(input any) (any, error) {
	logrus.Tracef("JsonToYamlConverter: Apply called with input %v", input)
	in_data, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid input type")
	}

	// Parse string to data
	var data any
	err := json.Unmarshal([]byte(in_data), &data)
	if err != nil {
		return nil, err
	}
	logrus.Tracef("JsonToYamlConverter: Unmarshalled JSON data: %v", data)

	// Use core processing logic
	result, err := j.ProcessData(data, j.config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ApplyStructured - structured execution for chain operations
func (j *JsonToYamlConverter) ApplyStructured(input *data.IntermediateData) (*data.IntermediateData, error) {
	logrus.Tracef("JsonToYamlConverter: ApplyStructured called with structured data")

	// Use core processing logic directly on structured data
	result, err := j.ProcessData(input.Data, j.config)
	if err != nil {
		return nil, err
	}

	// Convert result string back to structured data for YAML format
	// Since ProcessData returns a YAML string, we parse it back to structured data
	yamlData, err := data.NewFromYAML(result.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML result: %w", err)
	}

	return yamlData, nil
}

// ProcessData - core logic shared between both execution modes
func (j *JsonToYamlConverter) ProcessData(inputData any, config converters.BaseConfig) (any, error) {
	// Convert data to YAML
	out, err := mockableYamlMarshal(inputData)
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

// SupportsStructuredInput indicates this converter supports structured execution
func (j *JsonToYamlConverter) SupportsStructuredInput() bool {
	return true
}

// GetFormattingConfig returns no-op formatting config since YAML output doesn't use custom formatting
func (j *JsonToYamlConverter) GetFormattingConfig() converters.FormattingConfig {
	return converters.NoOpFormattingConfig{}
}

// Ensure JsonToYamlConverter implements all required interfaces
var _ converters.BaseConverter = (*JsonToYamlConverter)(nil)
var _ converters.StructuredConverter = (*JsonToYamlConverter)(nil)
var _ converters.ConfigurableConverter = (*JsonToYamlConverter)(nil)
var _ converters.DualModeConverter = (*JsonToYamlConverter)(nil)
