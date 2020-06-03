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
	"path"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var mergeYamlCmd = &cobra.Command{
	Use:   "env-to-yaml <.env-file-path> <yaml-file-path>",
	Short: "Add entries to a YAML element using a .env file as input",
	Aliases: []string{
		"envtoyaml",
		"env2yaml",
	},
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		envpath := args[0]
		yamlpath := args[1]

		dir, err := cmd.Flags().GetString("directory")
		handleError(err)

		ypath, err := cmd.Flags().GetString("ypath")
		handleError(err)

		err = lib.EnvToYamlFile(path.Join(dir, envpath), path.Join(dir, yamlpath), ypath)
		handleError(err)
	},
}

func init() {
	rootCmd.AddCommand(mergeYamlCmd)
	mergeYamlCmd.Flags().String("ypath", "spec.template.spec.containers.0.env", "A period separated string indicating where in the YAML the variables should be appended")
	mergeYamlCmd.Flags().StringP("directory", "d", "", "Directory where the files will be looked for")

	// mergeYamlCmd.Flags().StringP("name-key", "N", "name", "Specifies the key of the name of each entry")
	// mergeYamlCmd.Flags().StringP("value-key", "V", "value", "Specifies the key of the name of each entry")
}
