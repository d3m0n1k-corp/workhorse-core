package yaml_to_json

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"

	"gopkg.in/yaml.v3"
)

var mockableJsonMarshalIndent = json.MarshalIndent

type YamlToJsonConverter struct {
	config YamlToJsonConfig
}

// Apply - string-based execution for single operations
func (y *YamlToJsonConverter) Apply(input any) (any, error) {
	in_data, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid input type")
	}

	// Parse string to data
	var data any
	err := yaml.Unmarshal([]byte(in_data), &data)
	if err != nil {
		return nil, err
	}

	// Use core processing logic
	result, err := y.ProcessData(data, y.config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ApplyStructured - structured execution for chain operations
func (y *YamlToJsonConverter) ApplyStructured(input *data.IntermediateData) (*data.IntermediateData, error) {
	// Use core processing logic directly on structured data
	result, err := y.ProcessData(input.Data, y.config)
	if err != nil {
		return nil, err
	}

	// Convert result string back to structured data for JSON format
	// Since ProcessData returns a JSON string, we parse it back to structured data
	jsonData, err := data.NewFromJSON(result.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON result: %w", err)
	}

	return jsonData, nil
}

// ProcessData - core logic shared between both execution modes
func (y *YamlToJsonConverter) ProcessData(inputData any, config converters.BaseConfig) (any, error) {
	yamlConfig := config.(YamlToJsonConfig)

	// Build indent string
	indent := data.GetIndentString(yamlConfig.IndentType, yamlConfig.IndentSize)

	// Convert data to JSON with formatting (maintaining mockable function for testing)
	out, err := mockableJsonMarshalIndent(inputData, "", indent)
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

// SupportsStructuredInput indicates this converter supports structured execution
func (y *YamlToJsonConverter) SupportsStructuredInput() bool {
	return true
}

// GetFormattingConfig returns the converter's formatting configuration
func (y *YamlToJsonConverter) GetFormattingConfig() converters.FormattingConfig {
	return y.config
}

// Ensure YamlToJsonConverter implements all required interfaces
var _ converters.BaseConverter = (*YamlToJsonConverter)(nil)
var _ converters.StructuredConverter = (*YamlToJsonConverter)(nil)
var _ converters.ConfigurableConverter = (*YamlToJsonConverter)(nil)
var _ converters.DualModeConverter = (*YamlToJsonConverter)(nil)
