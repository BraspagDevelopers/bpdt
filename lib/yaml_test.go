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

	yamlReader := strings.NewReader(`
items:
  - name: first item
  - name: name of item
    fields:
      - name: key_1
        value: value_1
`)

	var buffer bytes.Buffer
	err := EnvToYaml(yamlReader, envReader, &buffer, "items.1.fields")
	require.NoError(t, err, "could not merge yaml")

	var expected interface{}
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
        value: '4'
`), &expected)
	require.NoError(t, err, "could not unmarshall expected yaml")
	require.NotEmpty(t, expected)

	var actual interface{}
	err = yaml.Unmarshal(buffer.Bytes(), &actual)
	require.NoError(t, err, "could not unmarshall actual yaml")
	require.NotEmpty(t, actual)

	assert.EqualValues(t, expected, actual, "yaml should match")
}

func TestReferenceSecrets(t *testing.T) {

	const (
		secretKeyPrefix = "#{"
		secretKeySuffix = "}#"
	)

	r := strings.NewReader(`
items:
  - name: name of item
    fields:
      - name: var01
        value: value_1
      - name: var02
        value: '#{secret_good}#'
      - name: var03
        value: '#notsecret{secret_bad}#'
      - name: var04
        value: 'word #{secret_good}# word'
`)

	var buffer bytes.Buffer
	// secrets := secretsYaml{
	// 	name: secretName,
	// 	keys: map[string]string{
	// 		secretKey: "abcde",
	// 	},
	// }
	err := ReferenceSecrets(r, &buffer, "items.0.fields", "my_secret_01", secretKeyPrefix, secretKeySuffix)
	require.NoError(t, err, "could not reference secrets in yaml")

	var expected interface{}
	err = yaml.Unmarshal([]byte(`
items:
  - name: name of item
    fields:
      - name: var01
        value: value_1
      - name: var02
        valueFrom:
          secretKeyRef:
            name: my_secret_01
            key: secret_good
      - name: var03
        value: '#notsecret{secret_bad}#'
      - name: var04
        value: 'word #{secret_good}# word'
`), &expected)
	require.NoError(t, err, "could not unmarshall expected yaml")
	require.NotEmpty(t, expected)

	var actual interface{}
	err = yaml.Unmarshal(buffer.Bytes(), &actual)
	require.NoError(t, err, "could not unmarshall actual yaml")
	require.NotEmpty(t, actual)
	assert.EqualValues(t, expected, actual, "yaml should match")
}

func TestLoadSecretsFromYaml(t *testing.T) {
	r := strings.NewReader(`
metadata:
  name: secretName0001
data:
  key1: value1
  key2: value2
  key3: 3
`)

	actual, err := LoadSecretsYaml(r)
	require.NoError(t, err)
	assert.EqualValues(t, "secretName0001", actual.name)
	expectedKeys := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "3",
	}
	assert.EqualValues(t, expectedKeys, actual.keys)
}
