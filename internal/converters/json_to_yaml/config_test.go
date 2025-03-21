package json_to_yaml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	config := JsonToYamlConfig{}
	require.NoError(t, config.Validate())
}
