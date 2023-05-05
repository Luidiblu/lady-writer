package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func executeCode(code string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	url := "https://api.openai.com/v1/chat/completions"

	message := Message{
		Role:    "system",
		Content: fmt.Sprintf("You are a helpful assistant that can execute code. Execute the following Golang code snippet: %s", code),
	}

	chatCompletionRequest := ChatCompletionRequest{
		Model:    "text-davinci-002",
		Messages: []Message{message},
	}

	reqBody, err := json.Marshal(chatCompletionRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error executing code: %s", string(respBody))
	}

	var chatCompletionResponse ChatCompletionResponse
	err = json.Unmarshal(respBody, &chatCompletionResponse)
	if err != nil {
		return "", err
	}

	response := chatCompletionResponse.Choices[0].Message.Content

	return response, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lady-writer <code>")
		os.Exit(1)
	}

	code := os.Args[1]

	result, err := executeCode(code)
	if err != nil {
		fmt.Printf("Error executing code: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %s\n", result)
}
