package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExport(t *testing.T) {
	config, err := New().
		AddJsonFile("file1.json").
		AddJsonFile("file2.json").
		Build()

	require.NoError(t, err, "Failed to build config")

	expected := map[string]string{
		"String01":       "new_sv01",
		"String02":       "new_sv02",
		"Array01__0":     "av03",
		"Array01__1":     "av04",
		"Object01__key1": "value1",
		"Object01__key2": "new_value2",
		"Object01__key3": "value3",
		"Object01__key4": "new_value4",
	}
	assert.EqualValues(t, expected, config)
}
