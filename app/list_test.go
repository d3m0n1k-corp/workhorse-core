package app

import (
	"reflect"
	"testing"
	"workhorse-core/internal/converters"

	"github.com/stretchr/testify/require"
)

func TestExtractConfTypes(t *testing.T) {
	object := struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}{"name", "type"}

	confs := extractConfTypes(reflect.TypeOf(object))
	require.Len(t, confs, 2)
	require.Equal(t, "name", confs[0].Name)
	require.Equal(t, "string", confs[0].Type)

}

func TestListConverters(t *testing.T) {

	mockableListConverters = func() []*converters.Registration {
		return []*converters.Registration{
			{
				Name:        "name",
				DemoInput:   1,
				Description: "description",
				InputType:   "input_type",
				OutputType:  "output_type",
				Config: reflect.TypeOf(struct {
					Name string `json:"name"`
					Type string `json:"type"`
				}{},
				),
			},
		}
	}

	reg_list := ListConverters()
	require.Len(t, reg_list, 1)
}
