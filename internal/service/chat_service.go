package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gogroq/internal/dto"
	"io"
	"net/http"
)

const (
	ApiBaseUrl = "https://api.groq.com/openai"
	SYSTEM     = "system"
	USER       = "user"

	LLMModelLlama38b       = "llama3-8b-8192"
	LLMModelLlama370b      = "llama3-70b-8192"
	LLMModelMixtral8x7b32k = "mixtral-8x7b-32768"
	LLMModelGemma7b        = "gemma-7b-it"
)

type GroqClient struct {
	ApiKey string `json:"api_url"`
}

func (groqClient *GroqClient) SendChatRequet(llmModel string, systemPrompt string, userPrompt string) (*dto.GroqChatResponse, error) {
	llm := llmModel

	if llmModel == "" {
		llm = LLMModelLlama38b
	}

	qroqMessages := make([]dto.GroqMessage, 0)

	if systemPrompt != "" {
		systemMessage := dto.GroqMessage{
			Role:    SYSTEM,
			Content: systemPrompt,
		}

		qroqMessages = append(qroqMessages, systemMessage)
	}

	if userPrompt == "" {
		err := errors.New("prompt is required")
		panic(err)
	}

	userMessage := dto.GroqMessage{
		Role:    USER,
		Content: userPrompt,
	}
	qroqMessages = append(qroqMessages, userMessage)

	chatRequest := dto.GroqChatRequest{
		Messages:    qroqMessages,
		Model:       llm,
		Temperature: 0,
		MaxTokens:   1024,
		TopP:        1,
		Stream:      false,
		Stop:        nil,
	}

	chatRequestJson, err := json.Marshal(chatRequest)
	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf("%s%s", ApiBaseUrl, "/v1/chat/completions")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(chatRequestJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", groqClient.ApiKey))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, reason: %s", res.StatusCode, res.Status)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	chatResponse := &dto.GroqChatResponse{}
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return nil, err
	}

	return chatResponse, err
}
