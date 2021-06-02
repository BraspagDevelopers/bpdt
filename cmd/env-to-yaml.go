package cmd

import (
	"path"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var envToYamlCmd = &cobra.Command{
	Use:   "env-to-yaml <.env-file-path> <yaml-file-path>",
	Short: "Add entries to a YAML element using a .env file as input",
	Aliases: []string{
		"envtoyaml",
		"env2yaml",
		"e2y",
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
	rootCmd.AddCommand(envToYamlCmd)
	envToYamlCmd.Flags().String("ypath", "spec.template.spec.containers.0.env", "A period separated string indicating where in the YAML the variables should be appended")
	envToYamlCmd.Flags().StringP("directory", "d", "", "Directory where the files will be looked for")

	// mergeYamlCmd.Flags().StringP("name-key", "N", "name", "Specifies the key of the name of each entry")
	// mergeYamlCmd.Flags().StringP("value-key", "V", "value", "Specifies the key of the name of each entry")
}
