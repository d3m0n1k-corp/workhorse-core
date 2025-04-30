package converter_list_test

import (
	"fmt"
	"os"
	"testing"
	"workhorse-core/app"

	"github.com/stretchr/testify/require"
)

func TestListConverters(t *testing.T) {

	//walk the internal/converters package and get the number of folders
	prefix := "../.."
	directory := "internal/converters"
	files, err := os.ReadDir(fmt.Sprintf("%s/%s", prefix, directory))
	folders := 0
	for _, file := range files {
		if file.IsDir() {
			folders++
		}
	}
	if err != nil {
		require.NoError(t, err, "Failed to read directory")
	}
	result := app.ListConverters()
	require.NotNil(t, result, "ListConverters should not return nil")
	require.NotEmpty(t, result, "ListConverters should not return an empty list")
	t.Logf("%d folders vs %d converters", folders, len(result))
	require.Len(t, result, folders, "ListConverters should return the same number of folders as in the internal/converters package")
}
