package json_prettifier

import (
	"encoding/json"
	"strings"
	"workhorse-core/internal/common/types"
)

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

	pretty_json, err := json.MarshalIndent(inp_json, "", indent)
	if err != nil {
		return nil, err
	}
	return pretty_json, nil
}
