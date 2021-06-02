package lib

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/palantir/stacktrace"
	"gopkg.in/yaml.v3"
)

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

	section, err := yamlGetByPath(doc, ypath)
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

	err = yamlSetByPath(doc, ypath, sectionMap)

	encoder := yaml.NewEncoder(w)
	err = encoder.Encode(&doc)
	if err != nil {
		return stacktrace.Propagate(err, "Could not encode yaml")
	}

	return nil
}

// EnvToYamlFile fill some yaml node with entries from and .env-formatted file.
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

// ReferenceSecrets
func ReferenceSecrets(r io.Reader, w io.Writer, ypath, secretname, prefix, suffix string) error {
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
	node, err := yamlGetByPath(doc, ypath)
	var data []T
	err = mapstructure.Decode(node, &data)
	if err != nil {
		return stacktrace.Propagate(err, "could not convert structure")
	}

	for i, item := range data {
		p := regexp.MustCompile(fmt.Sprintf("^%s(.*)%s$", prefix, suffix))
		matches := p.FindStringSubmatch(item.Value)
		if matches != nil {
			item.Value = ""
			item.ValueFrom.SecretKeyRef.Name = secretname
			item.ValueFrom.SecretKeyRef.Key = matches[1]

			data[i] = item
		}
	}
	e := yaml.NewEncoder(w)
	err = yamlSetByPath(doc, ypath, data)
	if err != nil {
		return stacktrace.Propagate(err, "could set the new value")
	}

	err = e.Encode(doc)
	if err != nil {
		return stacktrace.Propagate(err, "could not serialize yaml")
	}
	return nil
}

func ReferenceSecretsFile(path, ypath, secretname, prefix, suffix string) error {
	return readWrite(path, func(reader io.Reader, writer io.Writer) error {
		return ReferenceSecrets(reader, writer, ypath, secretname, prefix, suffix)
	})
}
