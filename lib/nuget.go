package lib

import (
	"bytes"
	"io"
	"os"

	"github.com/beevik/etree"
	"github.com/palantir/stacktrace"
)

// PatchNuget adds clear text password to a nuget config
func PatchNuget(reader io.Reader, writer io.Writer, source, username, password string) error {
	doc := etree.NewDocument()
	_, err := doc.ReadFrom(reader)
	if err != nil {
		return stacktrace.Propagate(err, "Could not read XML")
	}
	creds := doc.Root().
		CreateElement("packageSourceCredentials").
		CreateElement(source)
	add := creds.CreateElement("add")
	add.CreateAttr("key", "Username")
	add.CreateAttr("value", username)

	add = creds.CreateElement("add")
	add.CreateAttr("key", "ClearTextPassword")
	add.CreateAttr("value", password)

	doc.IndentTabs()
	_, err = doc.WriteTo(writer)
	if err != nil {
		return stacktrace.Propagate(err, "Could not write XML")
	}
	return nil
}

// PatchNugetFile adds clear text password to a nuget config file
func PatchNugetFile(path, source, username, password string) error {
	return openReadWrite(path, func(reader io.Reader, writer io.Writer) error {
		return PatchNuget(reader, writer, source, username, password)
	})
}

func openReadWrite(path string, function func(io.Reader, io.Writer) error) error {
	reader, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return stacktrace.Propagate(err, "Error opening file for reading")
	}
	defer reader.Close()

	var buffer bytes.Buffer
	err = function(reader, &buffer)
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
