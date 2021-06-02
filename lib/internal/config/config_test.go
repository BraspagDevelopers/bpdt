package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExport(t *testing.T) {
	config := Configuration{}

	file1, err := os.Open("file1.json")
	require.NoError(t, err, "Could not open file1.json")
	defer file1.Close()

	file2, err := os.Open("file2.json")
	require.NoError(t, err, "Could not open file2.json")
	defer file1.Close()

	err = config.AddJsonReader(file1)
	require.NoError(t, err, "Failed to add file1")

	err = config.AddJsonReader(file2)
	require.NoError(t, err, "Failed to add file2")

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
