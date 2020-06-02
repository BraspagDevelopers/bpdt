/*
Copyright © 2020 BPDT Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var exportEnvCmd = &cobra.Command{
	Use:     "export-settings",
	Aliases: []string{"ee"},
	Run: func(cmd *cobra.Command, args []string) {
		builder := lib.New()
		dir, err := cmd.Flags().GetString("directory")
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}

		filenames, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(2)
		}

		for _, jsonFileName := range filenames {
			builder = builder.AddJsonFile(path.Join(dir, jsonFileName))
		}

		config, err := builder.Build()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(3)
		}
		for k, v := range config {
			fmt.Println(fmt.Sprintf("%s=%s", k, v))
		}
	},
}

func init() {
	rootCmd.AddCommand(exportEnvCmd)
	exportEnvCmd.Flags().StringP("directory", "d", ".", "Determines where the files will be searched for")
	exportEnvCmd.Flags().StringArrayP("file", "f", []string{}, "Determines the files that will be used as input")
}
