package lib

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"gotest.tools/assert"
)

func TestGenerateConfigMap_ShouldNotFail(t *testing.T) {
	name := gofakeit.DomainName()
	key1 := gofakeit.DomainName()
	key2 := gofakeit.DomainName()
	key3 := gofakeit.DomainName()
	value1 := gofakeit.DomainName()
	value2 := gofakeit.DomainName()
	value3 := gofakeit.DomainName()

	data := map[string]string{
		key1: value1,
		key2: value2,
		key3: value3,
	}
	manifest, err := generateConfigMap(name, data)
	require.NoError(t, err)

	type ConfigMapYaml struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Name string `yaml:"name"`
		} `yaml:"metadata"`
		Data map[string]string `yaml:"data"`
	}
	var actual ConfigMapYaml
	err = yaml.Unmarshal([]byte(manifest), &actual)
	require.NoError(t, err)

	assert.Equal(t, "v1", actual.APIVersion)
	assert.Equal(t, "ConfigMap", actual.Kind)
	assert.DeepEqual(t, data, actual.Data)
}
