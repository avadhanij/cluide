package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/avadhanij/cluide/pkg/utils"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var model string 

type ChatRequest struct {
    Model    string `json:"model"`
	Store    bool  `json:"store"`
    Messages []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

var askChatCmd = &cobra.Command{
    Use:   "ask-chat",
    Short: "This subcommand directs queries to chatgpt. Post the query as an argument wrapped in quotes.",
    Long:  `Interacting with OpenAI's ChatGPT models require a API key that needs to be set as either as an environment
	variable - OPENAI_API_KEY, or provided in the cluide-config YAML file.`,
	Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		var apiToken string
		query := args[0]
		client := resty.New()
		viper.AutomaticEnv()
		viper.BindEnv("openai.api_key", "OPENAI_API_KEY")
		configPath := fmt.Sprintf("%s/.cluide-config.yml", os.Getenv("HOME"))

		if _, err := os.Stat(configPath); err == nil {
			viper.SetConfigName(".cluide-config")
			viper.SetConfigType("yml")
			viper.AddConfigPath("$HOME/")
			if err := viper.ReadInConfig(); err != nil {
				fmt.Printf("Error reading config file - %s\n", err)
			}
		}

		apiToken = viper.GetString("openai.api_key")
		if apiToken == "" {
			fmt.Println("Please set the OPENAI_API_KEY environment variable or provide it as part of cluide-config.")
			return
		}

		data := ChatRequest{
            Model: model,
			Store: false,
            Messages: []Message{
                {
                    Role:    "user",
                    Content: query,
                },
            },
        }
		jsonData, err := utils.CreateJSONString(data)
		if err != nil {
			fmt.Println("Error formatting JSON data:", err)
			return
		}

		resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(apiToken).
		SetBody(jsonData).
		Post("https://api.openai.com/v1/chat/completions")

		if err != nil {
			utils.PrintRespErr(err, resp)
			return
		} else {
			var respData map[string]any
			json.Unmarshal([]byte(resp.Body()), &respData)
			if checkError(respData) {
				return
			}
			choices := respData["choices"].([]interface{})

			for _, choice := range choices {
				choiceMap := choice.(map[string]interface{})
				message := choiceMap["message"].(map[string]interface{})
				color.Green("\n%s \n", message["content"].(string))
				break
			}
		}
    },
}

func checkError(response map[string]any) bool {
	error, ok := response["error"].(map[string]any)
	if ok {
		fmt.Println("Error: ", error["message"].(string))
		fmt.Println("Code: ", error["code"].(string))
		fmt.Println("Type: ", error["type"].(string))
	}
	return ok
}

func init() {
    rootCmd.AddCommand(askChatCmd)

	askChatCmd.Flags().StringVar(&model, "model", "gpt-4o-mini", "The model to use")
}