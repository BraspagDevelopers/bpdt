package cmd

import (
	"fmt"

	"github.com/BraspagDevelopers/bpdt/lib"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
}

var generateConfigMapCmd = &cobra.Command{
	Use:   "configmap <name-on-manifest>",
	Short: "Generate a ConfigMap manifest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fromEnv, err := cmd.Flags().GetBool("env")
		handleError(err)

		prefix, err := cmd.Flags().GetString("prefix")
		handleError(err)

		stripPrefix, err := cmd.Flags().GetBool("strip-prefix")
		handleError(err)

		ignoreCase, err := cmd.Flags().GetBool("ignore-case")
		handleError(err)

		configMap, err := lib.GenerateConfigMap(lib.GenerateConfigMapParams{
			Name:            name,
			FromEnvironment: fromEnv,
			Prefix:          prefix,
			StripPrefix:     stripPrefix,
			IgnoreCase:      ignoreCase,
		})
		handleError(err)

		fmt.Println(configMap)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateConfigMapCmd)

	generateConfigMapCmd.Flags().Bool("env", false, "Load variables from environment")
	generateConfigMapCmd.Flags().String("prefix", "", "Filter the variables by this prefix")
	generateConfigMapCmd.Flags().Bool("strip-prefix", false, "Strip the variable name prefix")
	generateConfigMapCmd.Flags().BoolP("ignore-case", "i", false, "Ignore  case  distinctions when filtering variable names")
}
