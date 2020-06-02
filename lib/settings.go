package lib

import (
	"bytes"
	"fmt"
	"io"

	"github.com/BraspagDevelopers/bpdt/lib/config"
	"github.com/palantir/stacktrace"
)

// ExportSettings exports settings
func ExportSettings(inputs []io.Reader, output io.Writer) error {
	config := config.Configuration{}
	for _, input := range inputs {
		err := config.AddJsonReader(input)
		if err != nil {
			return stacktrace.Propagate(err, "Could read configuration from %s", input)
		}
	}
	var buffer bytes.Buffer
	for k, v := range config {
		fmt.Fprintln(&buffer, fmt.Sprintf("%s=%s", k, v))
	}
	output.Write(buffer.Bytes())
	return nil
}
