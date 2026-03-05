package yaml_to_json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	obj := YamlToJsonConfig{}
	require.NoError(t, obj.Validate())
}
