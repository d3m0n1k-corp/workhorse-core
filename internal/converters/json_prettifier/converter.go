package json_prettifier

import (
	"encoding/json"
	"reflect"
	"strings"
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

func (j *JsonPrettifier) Apply(input any) (any, error) {
	inp_str := input.([]byte)
	var inp_json any
	err := json.Unmarshal(inp_str, &inp_json)
	if err != nil {
		return nil, err
	}

	var indent string
	if j.config.IndentType == "space" {
		indent = strings.Repeat(" ", j.config.IndentSize)
	} else {
		indent = strings.Repeat("\t", j.config.IndentSize)
	}

	pretty_json, err := mockableJsonMarshalIndent(inp_json, "", indent)
	if err != nil {
		return nil, err
	}
	return pretty_json, nil
}

var _ = converters.Register(&converters.Registration{
	Name:      "json_prettifier",
	DemoInput: []byte(`{"a":1,"b":2}`),
	Description: `
JsonPrettifier is a converter that takes a JSON input and returns a pretty-printed JSON output.
`,
	Config:     reflect.TypeOf(JsonPrettifierConfig{}),
	InputType:  types.JSON,
	OutputType: types.JSON,
})
