package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "cluide",
    Short: "Command line utility to send queries to popular AI platforms such as OpenAI's GPT-3 or Anthropic's Claude, etc. and receive responses.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use the --help flag to see available commands.")
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        return
    }
}

// Initialize the root command
func init() {
    // Here you can define flags and configuration settings.
}