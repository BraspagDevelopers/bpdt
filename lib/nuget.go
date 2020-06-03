package lib

import (
	"io"

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
	return readWrite(path, func(reader io.Reader, writer io.Writer) error {
		return PatchNuget(reader, writer, source, username, password)
	})
}
