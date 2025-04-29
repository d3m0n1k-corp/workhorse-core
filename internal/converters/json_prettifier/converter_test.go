package json_prettifier

import (
	"fmt"
	"testing"
	"workhorse-core/internal/common/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputType_returnType(t *testing.T) {
	j := JsonPrettifier{}
	assert.Equal(t, types.JSON, j.InputType())
}

func TestOutputType_returnType(t *testing.T) {
	j := JsonPrettifier{}
	assert.Equal(t, types.JSON, j.OutputType())
}

func TestApply_whenValidSpaces_returnPrettyJson(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 4, IndentType: "space"}}
	input := `{"a":1,"b":2}`
	output, err := j.Apply(input)
	require.NoError(t, err)
	expected := "{\n    \"a\": 1,\n    \"b\": 2\n}"
	assert.Equal(t, expected, output)
}

func TestApply_whenInvalidInput_returnError(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 4, IndentType: "space"}}
	input := `{"a":1,"b":2`
	_, err := j.Apply(input)
	assert.Error(t, err)
}

func TestApply_whenValidTabs_returnPrettyJson(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 1, IndentType: "tab"}}
	input := `{"a":1,"b":2}`
	output, err := j.Apply(input)
	require.NoError(t, err)
	expected := "{\n\t\"a\": 1,\n\t\"b\": 2\n}"
	require.Equal(t, expected, output)
}

func TestApply_whenMarshalError_returnError(t *testing.T) {
	mockableJsonMarshalIndent = func(v interface{}, prefix, indent string) ([]byte, error) {
		return nil, fmt.Errorf("Marshal error")
	}
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 1, IndentType: "tab"}}
	input := `{"a":1,"b":2}`
	_, err := j.Apply(input)
	require.Error(t, err)
}
