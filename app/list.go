package app

import (
	"fmt"
	"reflect"
	"workhorse-core/internal/converters"
)

var mockableListConverters = converters.ListConverters

type ItemConfig struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type RegisteredItem struct {
	Name        string        `json:"name"`
	DemoInput   string        `json:"demo_input"`
	Description string        `json:"description"`
	InputType   string        `json:"input_type"`
	OutputType  string        `json:"output_type"`
	Config      []*ItemConfig `json:"config"`
}

func extractConfTypes(t reflect.Type) []*ItemConfig {
	numFields := t.NumField()
	var confs []*ItemConfig
	for i := range numFields {
		field := t.Field(i)
		conf := ItemConfig{
			Name: field.Tag.Get("json"),
			Type: field.Type.String(),
		}
		confs = append(confs, &conf)
	}
	return confs
}

func ListConverters() []*RegisteredItem {
	conv_list := mockableListConverters()
	var reg_list []*RegisteredItem
	for _, reg := range conv_list {

		conf := extractConfTypes(reg.Config)

		reg_list = append(reg_list, &RegisteredItem{
			Name:        reg.Name,
			DemoInput:   fmt.Sprint(reg.DemoInput),
			Description: reg.Description,
			InputType:   reg.InputType,
			OutputType:  reg.OutputType,
			Config:      conf,
		})
	}
	return reg_list
}
