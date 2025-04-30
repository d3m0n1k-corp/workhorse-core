package json_to_yaml

import (
	"testing"
	"workhorse-core/internal/common/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestInputType(t *testing.T) {
	converter := JsonToYamlConverter{}
	require.Equal(t, types.JSON, converter.InputType())
}

func TestOutputType(t *testing.T) {
	converter := JsonToYamlConverter{}
	require.Equal(t, types.YAML, converter.OutputType())
}

func TestApply_ifInputNotString_returnError(t *testing.T) {
	converter := JsonToYamlConverter{}
	_, err := converter.Apply(123)
	require.Error(t, err)
}

func TestApply_ifInputIsInvalidJson_returnError(t *testing.T) {
	converter := JsonToYamlConverter{}
	_, err := converter.Apply("invalid json")
	require.Error(t, err)
}

func TestApply_ifYamlMarshalFails_returnError(t *testing.T) {
	mockableYamlMarshal = func(any) ([]byte, error) {
		return nil, assert.AnError
	}
	defer func() {
		mockableYamlMarshal = yaml.Marshal
	}()

	converter := JsonToYamlConverter{}
	_, err := converter.Apply(`{"a":1}`)
	require.Error(t, err)
}

func TestApply_ifInputIsValidJson_returnYaml(t *testing.T) {
	converter := JsonToYamlConverter{}
	yaml, err := converter.Apply(`{"a":1,"b":2}`)
	require.NoError(t, err)
	require.Equal(t, "a: 1\nb: 2\n", yaml)
}
