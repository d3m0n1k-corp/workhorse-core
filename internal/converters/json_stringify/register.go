package json_stringify

import (
	"reflect"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"
)

var _ = converters.Register(&converters.Registration{
	Name:        "json_stringify",
	DemoInput:   []byte(`{"a":1,"b":2}`),
	Description: "JsonStringifier is a converter that takes a JSON input and returns a stringified JSON output.",
	Config:      reflect.TypeOf(JsonStringifierConfig{}),
	InputType:   types.JSON,
	OutputType:  types.JSON_STRINGIFIED,
	Constructor: func(config converters.BaseConfig) converters.BaseConverter {
		return &JsonStringifier{config: config.(JsonStringifierConfig)}
	},
})
