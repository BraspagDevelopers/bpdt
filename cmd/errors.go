package cmd

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
