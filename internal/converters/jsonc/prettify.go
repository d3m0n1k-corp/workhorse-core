package jsonc

import (
	"encoding/json"
	"workhorse-core/internal/common"
	"workhorse-core/internal/converters"
	"workhorse-core/internal/converters/basec"
)

var _ = converters.Register("json-prettify", &JsonPrettifier{})

type JsonPrettifierConfig struct {
	Indent string `json:"indent" validate:"required"`
	Prefix string `json:"prefix" validate:"required"`
}

type JsonPrettifier struct {
	input  any
	output any
	logs   []string
	config JsonPrettifierConfig
	basec.BaseConverter
}

func NewJsonPrettifier(input common.Data[common.DataType], config JsonPrettifierConfig) *JsonPrettifier {
	return &JsonPrettifier{
		input:  input,
		logs:   []string{},
		config: config,
	}
}

func (j *JsonPrettifier) Apply() (any, error) {
	val, err := json.MarshalIndent(j.input, j.config.Prefix, j.config.Indent)
	j.output = val
	if err != nil {
		j.logs = append(j.logs, err.Error())
		return nil, err
	}
	return j.output, nil
}
