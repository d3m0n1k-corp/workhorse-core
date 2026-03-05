package json_to_yaml

import (
	"reflect"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters/base"
)

var _ = base.Register(&base.Registration{
	Name:      "json_to_yaml",
	DemoInput: []byte(`{"a":1,"b":2}`),
	Description: `
JsonToYamlConverter is a converter that takes a JSON input and returns a YAML output.
`,
	Config:     reflect.TypeOf(JsonToYamlConfig{}),
	InputType:  types.JSON,
	OutputType: types.YAML,
	Constructor: func(config base.BaseConfig) base.BaseConverter {
		return &JsonToYamlConverter{config: *config.(*JsonToYamlConfig)}
	},
})
