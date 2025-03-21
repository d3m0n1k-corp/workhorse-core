package yaml_to_json

import (
	"testing"
	"workhorse-core/internal/common/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputType(t *testing.T) {
	y := YamlToJsonConverter{}
	require.Equal(t, types.YAML, y.InputType())
}

func TestOutputType(t *testing.T) {
	y := YamlToJsonConverter{}
	require.Equal(t, types.JSON, y.OutputType())
}

func TestApply_whenInputIsNotString_returnError(t *testing.T) {
	y := YamlToJsonConverter{}
	_, err := y.Apply(1)
	require.Error(t, err)
}

func TestApply_whenUnmarshalFails_returnError(t *testing.T) {
	y := YamlToJsonConverter{}
	_, err := y.Apply("{:}")
	require.Error(t, err)
}

func TestApply_whenJsonMarshalIndentFails_returnError(t *testing.T) {
	mockableJsonMarshalIndent = func(_ any, _, _ string) ([]byte, error) {
		return nil, assert.AnError
	}
	y := YamlToJsonConverter{}
	_, err := y.Apply("a: 1")
	require.Error(t, err)
}

func TestApply_whenValidInput_returnJsonString(t *testing.T) {
	mockableJsonMarshalIndent = func(_ any, _, _ string) ([]byte, error) {
		return []byte(`{"a": 1}`), nil
	}
	y := YamlToJsonConverter{
		config: YamlToJsonConfig{
			IndentType: "space",
			IndentSize: 2,
		},
	}
	out, err := y.Apply("a: 1")
	require.NoError(t, err)
	require.Equal(t, `{"a": 1}`, out)
}
