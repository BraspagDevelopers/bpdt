package lib

import (
	"bytes"
	"fmt"
	"io"

	"github.com/BraspagDevelopers/bpdt/lib/config"
	"github.com/joho/godotenv"
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

	str, err := godotenv.Marshal(config)
	if err != nil {
		return stacktrace.Propagate(err, "could not write error")
	}
	output.Write([]byte(str))
	return nil
}
