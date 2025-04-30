package json_stringify

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	config := JsonStringifierConfig{}
	require.NoError(t, config.Validate())
}
