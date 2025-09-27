package json_prettifier

import (
	"encoding/json"
	"fmt"
	"strings"
	"workhorse-core/internal/common/types"
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
	inp_str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("invalid input type: expected string, got %T", input)
	}
	var inp_json any
	err := json.Unmarshal([]byte(inp_str), &inp_json)
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
	return string(pretty_json), nil
}
