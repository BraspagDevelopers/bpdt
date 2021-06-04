package cmd

import (
	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var patchNugetCmd = &cobra.Command{
	Use:   "patch-nuget <path> <nugetSource> <username> <password>",
	Short: "Add clear text passwords to a nuget config file",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		source := args[1]
		username := args[2]
		password := args[3]

		err := lib.PatchNugetFile(path, source, username, password)
		handleError(err)
	},
}

func init() {
	rootCmd.AddCommand(patchNugetCmd)
}
