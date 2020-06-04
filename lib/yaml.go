package lib

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
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

	m, isMap := y.(map[string]interface{})
	l, isList := y.([]interface{})
	switch {
	case isMap:
		val, hasKey := m[key]
		if !hasKey {
			return nil, fmt.Errorf(key)
		}
		return getByPath(val, subkey)
	case isList:
		index, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("not an index: %s", parts[0])
		}
		val := l[index]
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

	m, isMap := y.(map[string]interface{})
	l, isList := y.([]interface{})
	switch {
	case isMap:
		if subkey != "" {
			val, hasKey := m[key]
			if !hasKey {
				return fmt.Errorf(key)
			}
			return setByPath(val, subkey, value)
		}
		m[key] = value
	case isList:
		index, err := strconv.Atoi(key)
		if err != nil {
			return fmt.Errorf("not an index: %s", parts[0])
		}
		if subkey != "" {
			val := l[index]
			return setByPath(val, subkey, value)
		}
		l[index] = value
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
func EnvToYamlFile(envFilename, deploymentFilename, ypath string) error {
	envReader, err := os.Open(envFilename)
	if err != nil {
		return stacktrace.Propagate(err, "Could not open env file")
	}

	return readWrite(deploymentFilename, func(r io.Reader, w io.Writer) error {
		return EnvToYaml(r, envReader, w, ypath)
	})
}
