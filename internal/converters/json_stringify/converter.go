package json_stringify

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"
)

var mockableJsonMarshal = json.Marshal

type JsonStringifier struct {
	config JsonStringifierConfig
}

func (j *JsonStringifier) InputType() string {
	return types.JSON
}

func (j *JsonStringifier) OutputType() string {
	return types.JSON_STRINGIFIED
}

// Apply - string-based execution for single operations
func (j *JsonStringifier) Apply(input any) (any, error) {
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

	return result, nil
}

// ApplyStructured - structured execution for chain operations
func (j *JsonStringifier) ApplyStructured(input *data.IntermediateData) (*data.IntermediateData, error) {
	// Use core processing logic directly on structured data
	result, err := j.ProcessData(input.Data, j.config)
	if err != nil {
		return nil, err
	}

	// Return as stringified JSON format
	return data.NewFromData(result, "json_stringified"), nil
}

// ProcessData - core logic shared between both execution modes
func (j *JsonStringifier) ProcessData(inputData any, config converters.BaseConfig) (any, error) {
	// First marshal to compact JSON
	jsonBytes, err := mockableJsonMarshal(inputData)
	if err != nil {
		return nil, err
	}

	// Then stringify the JSON
	stringifiedBytes, err := mockableJsonMarshal(string(jsonBytes))
	if err != nil {
		return nil, err
	}

	return string(stringifiedBytes), nil
}

// SupportsStructuredInput indicates this converter supports structured execution
func (j *JsonStringifier) SupportsStructuredInput() bool {
	return true
}

// GetFormattingConfig returns no-op formatting config since stringify doesn't format
func (j *JsonStringifier) GetFormattingConfig() converters.FormattingConfig {
	return converters.NoOpFormattingConfig{}
}

// Ensure JsonStringifier implements all required interfaces
var _ converters.BaseConverter = (*JsonStringifier)(nil)
var _ converters.StructuredConverter = (*JsonStringifier)(nil)
var _ converters.ConfigurableConverter = (*JsonStringifier)(nil)
var _ converters.DualModeConverter = (*JsonStringifier)(nil)
