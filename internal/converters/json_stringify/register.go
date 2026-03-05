package json_stringify

import (
	"reflect"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters/base"
)

var _ = base.Register(&base.Registration{
	Name:        "json_stringify",
	DemoInput:   []byte(`{"a":1,"b":2}`),
	Description: "JsonStringifier is a converter that takes a JSON input and returns a stringified JSON output.",
	Config:      reflect.TypeOf(JsonStringifierConfig{}),
	InputType:   types.JSON,
	OutputType:  types.STRING,
	Constructor: func(config base.BaseConfig) base.BaseConverter {
		return &JsonStringifier{config: *config.(*JsonStringifierConfig)}
	},
})
