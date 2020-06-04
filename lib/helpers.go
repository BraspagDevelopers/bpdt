package lib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/palantir/stacktrace"
)

func readWrite(path string, f func(io.Reader, io.Writer) error) error {
	reader, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return stacktrace.Propagate(err, "Error opening file for reading")
	}
	defer reader.Close()

	var buffer bytes.Buffer
	err = f(reader, &buffer)
	if err != nil {
		return err
	}

	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return stacktrace.Propagate(err, "Error opening file for writing")
	}
	defer writer.Close()

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return stacktrace.Propagate(err, "Error writing data on file")
	}
	return nil
}

func yamlGetByPath(y interface{}, path string) (interface{}, error) {
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
		return yamlGetByPath(val, subkey)
	case []interface{}:
		index, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("not an index: %s", parts[0])
		}
		val := item[index]
		return yamlGetByPath(val, subkey)
	default:
		return nil, fmt.Errorf("unknown type")
	}
}

func yamlSetByPath(y interface{}, path string, value interface{}) error {
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
			return yamlSetByPath(val, subkey, value)
		}
		item[key] = value
	case []interface{}:
		index, err := strconv.Atoi(key)
		if err != nil {
			return fmt.Errorf("not an index: %s", parts[0])
		}
		if subkey != "" {
			val := item[index]
			return yamlSetByPath(val, subkey, value)
		}
		item[index] = value
	default:
		return fmt.Errorf("unknown type")
	}
	return nil
}
