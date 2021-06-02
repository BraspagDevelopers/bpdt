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
