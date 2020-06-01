/*
Copyright © 2020 NAME HERE andregsilv@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/BraspagDevelopers/bpdt/configuration"
	"github.com/spf13/cobra"
)

var exportEnvCmd = &cobra.Command{
	Use:     "export-env",
	Aliases: []string{"ee"},
	Run: func(cmd *cobra.Command, args []string) {
		builder := configuration.New()
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
