package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// var versionString string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cluide",
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("Unable to determine version information.")
			return
		}

		if buildInfo.Main.Version != "" {
			fmt.Printf("Version: %s\n", buildInfo.Main.Version)
		} else {
			fmt.Println("Version: unknown")
		}
	},
}

func init() {
    rootCmd.AddCommand(versionCmd)
}
