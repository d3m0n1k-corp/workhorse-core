package json_stringify

import (
	"encoding/json"
	"testing"
	"workhorse-core/internal/common/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputType(t *testing.T) {
	JsonStringifier := JsonStringifier{}
	require.Equal(t, types.JSON, JsonStringifier.InputType())
}

func TestOutputType(t *testing.T) {
	JsonStringifier := JsonStringifier{}
	require.Equal(t, types.JSON_STRINGIFIED, JsonStringifier.OutputType())
}

func TestApply_whenInvalidInput_returnError(t *testing.T) {
	JsonStringifier := JsonStringifier{}
	_, err := JsonStringifier.Apply("")
	require.Error(t, err)
}

func TestApply_whenMarshalError_returnError(t *testing.T) {
	mockableJsonMarshal = func(v any) ([]byte, error) {
		return nil, assert.AnError
	}
	defer func() {
		mockableJsonMarshal = json.Marshal
	}()

	JsonStringifier := JsonStringifier{}
	_, err := JsonStringifier.Apply(`{"a":1}`)
	require.Error(t, err)
}

func TestApply_when2ndMarshalError_returnError(t *testing.T) {
	count := 0
	mockableJsonMarshal = func(v any) ([]byte, error) {
		if count == 0 {
			count++
			return []byte(`{"a":1}`), nil
		}
		return nil, assert.AnError
	}
	defer func() {
		mockableJsonMarshal = json.Marshal
	}()

	JsonStringifier := JsonStringifier{}
	_, err := JsonStringifier.Apply(`{"a":1`)
	require.Error(t, err)
}
