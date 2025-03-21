package json_prettifier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate_whenInvalidStruct_returnError(t *testing.T) {
	config := JsonPrettifierConfig{
		IndentSize: 2,
		IndentType: "",
	}
	err := config.Validate()
	require.Error(t, err)
}

func TestValidate_whenOddSpaces_returnError(t *testing.T) {
	config := JsonPrettifierConfig{
		IndentSize: 3,
		IndentType: "space",
	}
	err := config.Validate()
	require.Error(t, err)
}

func TestValidate_whenTabIsNotOne_returnError(t *testing.T) {
	config := JsonPrettifierConfig{
		IndentSize: 2,
		IndentType: "tab",
	}
	err := config.Validate()
	require.Error(t, err)
}

func TestValidate_whenValidConfig_returnNil(t *testing.T) {
	config := JsonPrettifierConfig{
		IndentSize: 2,
		IndentType: "space",
	}
	err := config.Validate()
	require.NoError(t, err)
}
