package lib

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/palantir/stacktrace"
	"gopkg.in/yaml.v3"
)

func getByPath(y interface{}, path string) (interface{}, error) {
	if path == "" {
		return y, nil
	}

	parts := strings.SplitN(path, ".", 2)
	key := parts[0]
	subkey := ""
	if len(parts) > 1 {
		subkey = parts[1]
	}

	switch item := y.(type) {
	case map[string]interface{}:
		val, hasKey := item[key]
		if !hasKey {
			return nil, fmt.Errorf(key)
		}
		return getByPath(val, subkey)
	case []interface{}:
		index, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("not an index: %s", parts[0])
		}
		val := item[index]
		return getByPath(val, subkey)
	default:
		return nil, fmt.Errorf("unknown type")
	}
}

func setByPath(y interface{}, path string, value interface{}) error {
	if path == "" {
		return fmt.Errorf("Could not set the root")
	}

	parts := strings.SplitN(path, ".", 2)
	key := parts[0]
	subkey := ""
	if len(parts) > 1 {
		subkey = parts[1]
	}

	switch item := y.(type) {
	case map[string]interface{}:
		if subkey != "" {
			val, hasKey := item[key]
			if !hasKey {
				return fmt.Errorf(key)
			}
			return setByPath(val, subkey, value)
		}
		item[key] = value
	case []interface{}:
		index, err := strconv.Atoi(key)
		if err != nil {
			return fmt.Errorf("not an index: %s", parts[0])
		}
		if subkey != "" {
			val := item[index]
			return setByPath(val, subkey, value)
		}
		item[index] = value
	default:
		return fmt.Errorf("unknown type")
	}
	return nil
}

// EnvToYaml fill some yaml node with entries from and .env-formatted file
func EnvToYaml(r, envr io.Reader, w io.Writer, ypath string) error {
	var doc interface{}
	decoder := yaml.NewDecoder(r)
	err := decoder.Decode(&doc)
	if err != nil {
		return stacktrace.Propagate(err, "could not decode yaml")
	}

	env, err := godotenv.Parse(envr)
	if err != nil {
		return stacktrace.Propagate(err, "could not parse env")
	}

	section, err := getByPath(doc, ypath)
	if err != nil {
		return stacktrace.Propagate(err, "could find path on yaml")
	}
	sectionMap, isMap := section.([]interface{})
	if !isMap {
		return fmt.Errorf("could not cast section")
	}

	for k, v := range env {
		sectionMap = append(sectionMap, map[string]string{
			"name":  k,
			"value": v,
		})
	}

	err = setByPath(doc, ypath, sectionMap)

	encoder := yaml.NewEncoder(w)
	err = encoder.Encode(&doc)
	if err != nil {
		return stacktrace.Propagate(err, "Could not encode yaml")
	}

	return nil
}

// EnvToYaml fill some yaml node with entries from and .env-formatted file.
// Reads and writes directly into files, instead of readers and writers.
func EnvToYamlFile(envFilename, yamlFilename, ypath string) error {
	envReader, err := os.Open(envFilename)
	if err != nil {
		return stacktrace.Propagate(err, "Could not open env file")
	}

	return readWrite(yamlFilename, func(r io.Reader, w io.Writer) error {
		return EnvToYaml(r, envReader, w, ypath)
	})
}

type secretsYaml struct {
	name string
	keys map[string]string
}

func ReferenceSecrets(r io.Reader, w io.Writer, ypath string, s secretsYaml) error {
	d := yaml.NewDecoder(r)
	var doc interface{}
	err := d.Decode(&doc)
	if err != nil {
		return stacktrace.Propagate(err, "could not decode yaml")
	}

	type T struct {
		Name      string
		Value     string `yaml:",omitempty"`
		ValueFrom struct {
			SecretKeyRef struct {
				Name string
				Key  string
			} `yaml:"secretKeyRef"`
		} `yaml:"valueFrom,omitempty"`
	}
	node, err := getByPath(doc, ypath)
	var data []T
	err = mapstructure.Decode(node, &data)
	if err != nil {
		return stacktrace.Propagate(err, "could not convert structure")
	}

	for i, item := range data {
		p := regexp.MustCompile("#<(.*)>#")
		matches := p.FindStringSubmatch(item.Value)
		if matches != nil {
			item.Value = ""
			item.ValueFrom.SecretKeyRef.Name = s.name
			item.ValueFrom.SecretKeyRef.Key = matches[1]

			data[i] = item
		}
	}
	e := yaml.NewEncoder(w)
	err = setByPath(doc, ypath, data)
	if err != nil {
		return stacktrace.Propagate(err, "could set the new value")
	}

	err = e.Encode(doc)
	if err != nil {
		return stacktrace.Propagate(err, "could not serialize yaml")
	}
	return nil
}

func LoadSecretsYaml(r io.Reader) (result secretsYaml, err error) {
	d := yaml.NewDecoder(r)
	var data struct {
		Metadata struct {
			Name string
		}
		Data map[string]string
	}
	err = d.Decode(&data)
	if err != nil {
		err = stacktrace.Propagate(err, "could not parse yaml")
		return
	}
	result = secretsYaml{
		name: data.Metadata.Name,
		keys: data.Data,
	}
	return
}
