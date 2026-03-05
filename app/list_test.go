package app

import (
	"reflect"
	"testing"
	"workhorse-core/internal/converters/base"

	"github.com/stretchr/testify/require"
)

func TestExtractConfTypes(t *testing.T) {
	object := struct {
		Name       string `json:"name" help:"The name of the item"`
		Type       string `json:"type" default:"default_type"`
		IndentType string `json:"indent_type" validate:"required,oneof=space tab" help:"Type of indentation"`
	}{"name", "type", "space"}

	confs := extractConfTypes(reflect.TypeOf(object))
	require.Len(t, confs, 3)

	// Test first field (Name)
	require.Equal(t, "name", confs[0].Name)
	require.Equal(t, "string", confs[0].Type)
	require.Equal(t, "The name of the item", confs[0].Help)
	require.Empty(t, confs[0].Default)
	require.Nil(t, confs[0].Options)

	// Test second field (Type)
	require.Equal(t, "type", confs[1].Name)
	require.Equal(t, "string", confs[1].Type)
	require.Equal(t, "default_type", confs[1].Default)
	require.Empty(t, confs[1].Help)
	require.Nil(t, confs[1].Options)

	// Test third field (IndentType) with options
	require.Equal(t, "indent_type", confs[2].Name)
	require.Equal(t, "string", confs[2].Type)
	require.Equal(t, "Type of indentation", confs[2].Help)
	require.Empty(t, confs[2].Default)
	require.Equal(t, []any{"space", "tab"}, confs[2].Options)
}

func TestListConverters(t *testing.T) {

	mockableListConverters = func() []*base.Registration {
		return []*base.Registration{
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
