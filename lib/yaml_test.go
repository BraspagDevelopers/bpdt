package lib

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestEnvToYaml(t *testing.T) {

	envReader := strings.NewReader(`
key_2=value_2
key_3=value 3 with spaces
key_4=4
`)

	deploymentReader := strings.NewReader(`
items:
  - name: first item
  - name: name of item
    fields:
      - name: key_1
        value: value_1
`)

	var buffer bytes.Buffer
	err := EnvToYaml(deploymentReader, envReader, &buffer, "items.1.fields")
	require.NoError(t, err, "could not merge yaml")

	type T struct {
		Items []struct {
			Name   string
			Fields []struct {
				Name  string
				Value string
			}
		}
	}
	var expected T
	err = yaml.Unmarshal([]byte(`
items:
  - name: first item
  - name: name of item
    fields:
      - name: key_1
        value: value_1
      - name: key_2
        value: value_2
      - name: key_3
        value: 'value 3 with spaces'
      - name: key_4
        value: 4
`), &expected)
	require.NoError(t, err, "could not unmarshall expected yaml")
	require.NotEmpty(t, expected)

	var actual T
	err = yaml.Unmarshal(buffer.Bytes(), &actual)
	require.NoError(t, err, "could not unmarshall actual yaml")
	require.NotEmpty(t, actual)

	assert.EqualValues(t, expected, actual, "yaml should match")
}
