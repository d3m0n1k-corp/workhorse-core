package json_prettifier

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"
)

var mockableJsonMarshalIndent = json.MarshalIndent

type JsonPrettifier struct {
	config JsonPrettifierConfig
}

func (j *JsonPrettifier) InputType() string {
	return types.JSON
}

func (j *JsonPrettifier) OutputType() string {
	return types.JSON
}

// Apply - string-based execution for single operations with object pooling optimization
func (j *JsonPrettifier) Apply(input any) (any, error) {
	inp_str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("invalid input type: expected string, got %T", input)
	}

	// Parse string to data
	var inp_json any
	err := json.Unmarshal([]byte(inp_str), &inp_json)
	if err != nil {
		return nil, err
	}

	// Use core processing logic
	result, err := j.ProcessData(inp_json, j.config)
	if err != nil {
		return nil, err
	}

	// Serialize back to string with formatting (maintaining mockable function for testing)
	indent := data.GetIndentString(j.config.IndentType, j.config.IndentSize)
	pretty_json, err := mockableJsonMarshalIndent(result, "", indent)
	if err != nil {
		return nil, err
	}
	return string(pretty_json), nil
} // ApplyStructured - structured execution for chain operations
func (j *JsonPrettifier) ApplyStructured(input *data.IntermediateData) (*data.IntermediateData, error) {
	// Use core processing logic directly on structured data
	result, err := j.ProcessData(input.Data, j.config)
	if err != nil {
		return nil, err
	}

	// Return structured data without serialization
	// The formatting will be applied when the data is finally serialized
	return data.NewFromData(result, "json"), nil
}

// ProcessData - core logic shared between both execution modes
func (j *JsonPrettifier) ProcessData(inputData any, config converters.BaseConfig) (any, error) {
	// For JSON prettifier, the core logic is just returning the same data
	// The formatting happens during serialization
	return inputData, nil
}

// SupportsStructuredInput indicates this converter supports structured execution
func (j *JsonPrettifier) SupportsStructuredInput() bool {
	return true
}

// GetFormattingConfig returns the converter's formatting configuration
func (j *JsonPrettifier) GetFormattingConfig() converters.FormattingConfig {
	return j.config
}

// Ensure JsonPrettifier implements all required interfaces
var _ converters.BaseConverter = (*JsonPrettifier)(nil)
var _ converters.StructuredConverter = (*JsonPrettifier)(nil)
var _ converters.ConfigurableConverter = (*JsonPrettifier)(nil)
var _ converters.DualModeConverter = (*JsonPrettifier)(nil)
