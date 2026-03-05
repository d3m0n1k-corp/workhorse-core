package app

import (
	"fmt"
	"reflect"
	"strings"
	"workhorse-core/internal/converters"
)

var mockableListConverters = converters.ListConverters

type ItemConfig struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Default string `json:"default,omitempty"`
	Help    string `json:"help,omitempty"`
	Options []any  `json:"options,omitempty"`
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
			Name:    field.Tag.Get("json"),
			Type:    field.Type.String(),
			Default: field.Tag.Get("default"),
			Help:    field.Tag.Get("help"),
			Options: extractOptionsFromValidate(field.Tag.Get("validate")),
		}
		confs = append(confs, &conf)
	}
	return confs
}

// extractOptionsFromValidate parses the validate tag and extracts options from oneof constraints
func extractOptionsFromValidate(validateTag string) []any {
	if validateTag == "" {
		return nil
	}

	// Split by comma to get individual validation rules
	for rule := range strings.SplitSeq(validateTag, ",") {
		rule = strings.TrimSpace(rule)

		// Look for oneof constraints
		if after, ok := strings.CutPrefix(rule, "oneof="); ok {
			optionsStr := after
			optionsList := strings.Fields(optionsStr) // Split by whitespace

			// Convert to []any
			var options []any
			for _, opt := range optionsList {
				options = append(options, opt)
			}
			return options
		}
	}

	return nil
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
