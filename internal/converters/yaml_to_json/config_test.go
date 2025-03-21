package yaml_to_json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate_whenInvalidStruct_returnError(t *testing.T) {
	obj := YamlToJsonConfig{
		IndentSize: 1,
		IndentType: "",
	}
	require.Error(t, obj.Validate())
}

func TestValidate_whenOddSpaces_returnError(t *testing.T) {
	obj := YamlToJsonConfig{
		IndentSize: 3,
		IndentType: "space",
	}
	require.Error(t, obj.Validate())
}

func TestValidate_whenTabSizeNotOne_returnError(t *testing.T) {
	obj := YamlToJsonConfig{
		IndentSize: 2,
		IndentType: "tab",
	}
	require.Error(t, obj.Validate())
}

func TestValidate_whenValidStruct_returnNil(t *testing.T) {
	obj := YamlToJsonConfig{
		IndentSize: 2,
		IndentType: "space",
	}
	require.Nil(t, obj.Validate())
}
