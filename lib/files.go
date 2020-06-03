package lib

import (
	"bytes"
	"io"
	"os"

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
