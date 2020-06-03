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
	"io"
	"os"
	"path"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var exportEnvCmd = &cobra.Command{
	Use:   "export-settings",
	Short: "Convert multiples `appsettings.*.json` files to .env file syntax",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString("directory")
		handleError(err)

		filenames, err := cmd.Flags().GetStringArray("file")
		handleError(err)

		readers := make([]io.Reader, len(filenames))
		for i, jsonFileName := range filenames {
			file, err := os.Open(path.Join(dir, jsonFileName))
			handleError(err)

			defer file.Close()
			readers[i] = file
		}

		lib.ExportSettings(readers, os.Stdout)
		handleError(err)
	},
}

func init() {
	rootCmd.AddCommand(exportEnvCmd)
	exportEnvCmd.Flags().StringP("directory", "d", "", "Directory where the files will be looked for")
	exportEnvCmd.Flags().StringArrayP("file", "f", []string{
		"appsettings.json",
		"appsettings.Development.json",
	}, "Files that will be used as input")
}
