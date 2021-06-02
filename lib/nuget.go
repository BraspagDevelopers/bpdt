package lib

import (
	"io"

	"github.com/beevik/etree"
	"github.com/palantir/stacktrace"
)

func findOrCreateElement(xml *etree.Element, name string) *etree.Element {
	if el := xml.SelectElement(name); el != nil {
		return el
	}
	return xml.CreateElement(name)
}

func findOrCreateElementWithAttribute(xml *etree.Element, name, attrName, attrValue string) *etree.Element {
	if els := xml.SelectElements(name); len(els) > 0 {
		for _, el := range els {
			if attr := el.SelectAttr(attrName); attr != nil {
				if attr.Value == attrValue {
					return el
				}
			}
		}
	}
	el := xml.CreateElement(name)
	el.CreateAttr(attrName, attrValue)
	return el
}

func setAttribute(xml *etree.Element, name string, value string) *etree.Attr {
	if attr := xml.SelectAttr(name); attr != nil {
		attr.Value = value
		return attr
	}
	return xml.CreateAttr(name, value)
}

// PatchNuget adds clear text password to a nuget config
func PatchNuget(reader io.Reader, writer io.Writer, source, username, password string) error {
	doc := etree.NewDocument()
	_, err := doc.ReadFrom(reader)
	if err != nil {
		return stacktrace.Propagate(err, "Could not read XML")
	}
	pkgSrcCreds := findOrCreateElement(doc.Root(), "packageSourceCredentials")
	creds := findOrCreateElement(pkgSrcCreds, source)

	add := findOrCreateElementWithAttribute(creds, "add", "key", "Username")
	setAttribute(add, "value", username)

	add = findOrCreateElementWithAttribute(creds, "add", "key", "ClearTextPassword")
	setAttribute(add, "value", password)

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
