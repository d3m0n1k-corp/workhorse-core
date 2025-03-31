package yaml_to_json

import (
	"reflect"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"
)

var _ = converters.Register(&converters.Registration{
	Name:      "yaml_to_json",
	DemoInput: []byte(`a: 1, b: 2`),
	Description: `
YamlToJsonConverter is a converter that takes a YAML input and returns a JSON output.
`,
	Config:     reflect.TypeOf(YamlToJsonConfig{}),
	InputType:  types.YAML,
	OutputType: types.JSON,
	Constructor: func(config converters.BaseConfig) converters.BaseConverter {
		return &YamlToJsonConverter{config: *config.(*YamlToJsonConfig)}
	},
})
