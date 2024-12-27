package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/avadhanij/cluide/pkg/utils"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
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
    Long:  `Interacting with OpenAI's ChatGPT models require a API key that needs to be set as an environment
	variable - OPENAI_API_KEY. The question should be passed as an argument to the subcommand.`,
	Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		client := resty.New()
		apiToken := os.Getenv("OPENAI_API_KEY")

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