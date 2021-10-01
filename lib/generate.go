package lib

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type GenerateConfigMapParams struct {
	Name            string
	FromEnvironment bool
	Prefix          string
	StripPrefix     bool
}

func (p GenerateConfigMapParams) Validate() error {
	if !p.FromEnvironment {
		return errors.New("environment is the only source for now but was not specified")
	}
	return nil
}

func (p GenerateConfigMapParams) getVariables() (map[string]string, error) {
	envLines := os.Environ()
	result := make(map[string]string, 0)
	for n, envLine := range envLines {
		splitted := strings.SplitN(envLine, "=", 2)
		if len(splitted) != 2 {
			return nil, fmt.Errorf("the %dth environment variable and value could not be parsed", n)
		}

		key := splitted[0]
		value := splitted[1]

		if p.isValidKey(key) {
			key = p.normalizeKey(key)
			result[key] = value
		}
	}

	return result, nil
}

func (p GenerateConfigMapParams) isValidKey(key string) bool {
	return strings.HasPrefix(key, p.Prefix)
}

func (p GenerateConfigMapParams) normalizeKey(key string) string {
	if p.StripPrefix {
		key = key[len(p.Prefix):]
	}
	return key
}

func (p GenerateConfigMapParams) marshalConfigMap(dataMap map[string]string) (string, error) {
	manifest := yaml.MapSlice{
		{"apiVersion", "v1"},
		{"kind", "ConfigMap"},
		{"metadata", map[string]string{
			"name": p.Name,
		}},
		{"data", dataMap},
	}
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GenerateConfigMap(params GenerateConfigMapParams) (string, error) {
	err := params.Validate()
	if err != nil {
		return "", err
	}

	envVars, err := params.getVariables()
	if err != nil {
		return "", err
	}

	return params.marshalConfigMap(envVars)
}
