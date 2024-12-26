package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// FormatOutput formats the output as a JSON string.
func FormatOutput(data interface{}) (string, error) {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func PrintRespErr(err error, resp *resty.Response) {
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Proto      :", resp.Proto())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()
}

// ParseInput reads input from the user and returns it as a string.
func ParseInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}