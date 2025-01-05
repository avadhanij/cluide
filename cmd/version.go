package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionString string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cluide",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cluide %s\n", versionString)
	},
}

func init() {
    rootCmd.AddCommand(versionCmd)
}