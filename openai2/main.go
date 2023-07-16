package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float32   `json:"temperature"`
	MaxTokens        int       `json:"max_tokens"`
	TopP             float32   `json:"top_p"`
	FrequencyPenalty float32   `json:"frequency_penalty"`
	PresencePenalty  float32   `json:"presence_penalty"`
}

func chatCompletions() {
	// Define the API endpoint
	url := "https://api.openai.com/v1/chat/completions"

	// Create the request payload
	messages := []Message{
		{
			"system",
			"You will be provided with statements, and your task is to convert them to standard English.",
		},
		{
			"user",
			"She no went to the market.",
		},
	}
	payload := Payload{
		"gpt-3.5-turbo",
		messages,
		0,
		64,
		1.0,
		0.0,
		0.0,
	}

	// Convert the payload to JSON
	jsonString, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Add the necessary headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	chatCompletions()
}
