package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "cluide",
    Short: "Command line buddy to send queries to AI platforms",
    Long:  `Cluide is a command line tool that allows you to send queries to AI platforms like OpenAI's GPT-3 and receive responses.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hi, I am Cluide! Use the --help flag to see available commands.")
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