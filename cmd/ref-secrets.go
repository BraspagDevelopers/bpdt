package cmd

import (
	"path"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

// linkSecretsCmd represents the linkSecrets command
var refSecretsCmd = &cobra.Command{
	Use: "ref-secrets <file-path> <secret-name>",
	Aliases: []string{
		"refsecrets",
		"refsec",
		"rs",
	},
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		secretname := args[1]

		ypath, err := cmd.Flags().GetString("ypath")
		handleError(err)

		prefix, err := cmd.Flags().GetString("prefix")
		handleError(err)

		suffix, err := cmd.Flags().GetString("suffix")
		handleError(err)

		directory, err := cmd.Flags().GetString("directory")
		handleError(err)

		lib.ReferenceSecretsFile(path.Join(directory, filename), ypath, secretname, prefix, suffix)
	},
}

func init() {
	rootCmd.AddCommand(refSecretsCmd)
	refSecretsCmd.Flags().String("ypath", "spec.template.spec.containers.0.env", "A period separated string indicating where in the YAML the variables are placed")
	refSecretsCmd.Flags().StringP("directory", "d", "", "Directory where the files will be looked for")

	refSecretsCmd.Flags().StringP("prefix", "p", "#<Secret>{", "The prefix for the secret variables")
	refSecretsCmd.Flags().StringP("suffix", "s", "}#", "The suffix for the secret variables")
}
