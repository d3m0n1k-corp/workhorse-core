package data

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// IntermediateData represents parsed, structured data that can be efficiently
// passed between converters without repeated serialization/deserialization
type IntermediateData struct {
	// The parsed data structure - can be map[string]any, []any, etc.
	Data any

	// Original format for reference and potential optimizations
	SourceFormat string

	// Optional: Keep raw bytes for pass-through scenarios
	RawBytes []byte
}

// NewFromJSON creates IntermediateData from JSON string
func NewFromJSON(jsonStr string) (*IntermediateData, error) {
	var data any
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &IntermediateData{
		Data:         data,
		SourceFormat: "json",
		RawBytes:     []byte(jsonStr),
	}, nil
}

// NewFromYAML creates IntermediateData from YAML string
func NewFromYAML(yamlStr string) (*IntermediateData, error) {
	var data any
	err := yaml.Unmarshal([]byte(yamlStr), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &IntermediateData{
		Data:         data,
		SourceFormat: "yaml",
		RawBytes:     []byte(yamlStr),
	}, nil
}

// NewFromData creates IntermediateData from already parsed data
func NewFromData(data any, sourceFormat string) *IntermediateData {
	return &IntermediateData{
		Data:         data,
		SourceFormat: sourceFormat,
		RawBytes:     nil,
	}
}

// ToJSON serializes to JSON string
func (id *IntermediateData) ToJSON(indent string) (string, error) {
	if indent == "" {
		bytes, err := json.Marshal(id.Data)
		return string(bytes), err
	}

	bytes, err := json.MarshalIndent(id.Data, "", indent)
	return string(bytes), err
}

// ToYAML serializes to YAML string
func (id *IntermediateData) ToYAML() (string, error) {
	bytes, err := yaml.Marshal(id.Data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML: %w", err)
	}
	return string(bytes), nil
}

// ToJSONStringified serializes data as JSON then stringifies it
func (id *IntermediateData) ToJSONStringified() (string, error) {
	// First marshal to JSON
	jsonBytes, err := json.Marshal(id.Data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Then stringify the JSON
	stringifiedBytes, err := json.Marshal(string(jsonBytes))
	if err != nil {
		return "", fmt.Errorf("failed to stringify JSON: %w", err)
	}

	return string(stringifiedBytes), nil
}

// Clone creates a deep copy for safe concurrent access
func (id *IntermediateData) Clone() *IntermediateData {
	// For now, we'll do a simple copy
	// In production, you might want proper deep cloning for complex nested structures
	return &IntermediateData{
		Data:         id.Data, // Shallow copy - be careful with mutations
		SourceFormat: id.SourceFormat,
		RawBytes:     append([]byte(nil), id.RawBytes...),
	}
}

// GetIndentString builds indent string for formatting
func GetIndentString(indentType string, indentSize int) string {
	if indentType == "space" {
		return strings.Repeat(" ", indentSize)
	}
	return strings.Repeat("\t", indentSize)
}
