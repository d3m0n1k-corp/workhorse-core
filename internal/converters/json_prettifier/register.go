package json_prettifier

import (
	"reflect"
	"workhorse-core/internal/common/types"
	"workhorse-core/internal/converters"
)

var _ = converters.Register(&converters.Registration{
	Name:        "json_prettifier",
	DemoInput:   []byte(`{"a":1,"b":2}`),
	Description: "JsonPrettifier is a converter that takes a JSON input and returns a pretty-printed JSON output.",
	Config:      reflect.TypeOf(JsonPrettifierConfig{}),
	InputType:   types.JSON,
	OutputType:  types.JSON,
	Constructor: func(config converters.BaseConfig) converters.BaseConverter {
		return &JsonPrettifier{config: *config.(*JsonPrettifierConfig)}
	},
})
