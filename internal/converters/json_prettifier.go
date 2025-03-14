package converters

import (
	"encoding/json"
	"strings"
	"workhorse-core/internal/common/types"

	"github.com/go-playground/validator/v10"
)

var vd = validator.New()

type JsonPrettifierConfig struct {
	IndentSize int    `json:"indent_size" validate:"required"`
	Prefix     string `json:"prefix" validate:"required"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab"`
}

func (j *JsonPrettifierConfig) Validate() error {
	err := vd.Struct(j)
	if err != nil {
		return err
	}
	return nil
}

//Sec

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

	pretty_json, err := json.MarshalIndent(inp_json, j.config.Prefix, indent)
	if err != nil {
		return nil, err
	}
	return pretty_json, nil
}
