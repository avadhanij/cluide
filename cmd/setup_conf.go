package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setupConfCmd = &cobra.Command{
	Use:   "setup-config",
	Short: "This subcommand sets up the configuration file for cluide.",
	Long:  `This subcommand sets up the configuration file for cluide under $HOME/.config/cluide. The configuration file is used to store API 
	keys and any settings associated with each AI platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFolder := fmt.Sprintf("%s/.config/cluide", os.Getenv("HOME"))
		configPath := fmt.Sprintf("%s/config.toml", configFolder)

		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("Config file already exists.")
			return
		}

		os.MkdirAll(configFolder, os.ModePerm)

		viper.SetDefault("openai.api_key", "<your_openai_api_key>")
		viper.SetDefault("anthropic.api_key", "<your_anthropic_api_key>")

		viper.SafeWriteConfigAs(configPath)
		fmt.Printf("Config file created at %s\n", configPath)
	},
}

func init() {
	rootCmd.AddCommand(setupConfCmd)
}