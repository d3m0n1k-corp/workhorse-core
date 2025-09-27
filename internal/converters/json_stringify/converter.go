package json_stringify

import (
	"encoding/json"
	"fmt"
	"workhorse-core/internal/common/types"
)

var mockableJsonMarshal = json.Marshal

type JsonStringifier struct {
	config JsonStringifierConfig
}

func (j *JsonStringifier) InputType() string {
	return types.JSON
}

func (j *JsonStringifier) OutputType() string {
	return types.JSON_STRINGIFIED
}

func (j *JsonStringifier) Apply(input any) (any, error) {
	inp_str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("invalid input type: expected string, got %T", input)
	}
	var inp_json any
	err := json.Unmarshal([]byte(inp_str), &inp_json)
	if err != nil {
		return nil, err
	}
	result, err := mockableJsonMarshal(inp_json)
	if err != nil {
		return nil, err
	}

	result, err = mockableJsonMarshal(string(result))
	if err != nil {
		return nil, err
	}

	return string(result), nil
}
