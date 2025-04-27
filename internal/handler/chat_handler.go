package handler

import (
	"encoding/json"
	"errors"
	"gogroq/internal/dto"
	"gogroq/internal/service"
	"net/http"
	"os"
)

const (
	SystemPrompt = "Você é um melhor amigo do usuário e o ajudará com o que ele precisar." +
		" Responda às perguntas no idioma que o usuário perguntou."

	UserPrompt = "%s"
)

func ChatPost(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		err := errors.New("UNAUTHORIZED")
		panic(err)
	}

	var qroqMessage dto.GroqMessage
	if err := json.NewDecoder(r.Body).Decode(&qroqMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groqClient := &service.GroqClient{ApiKey: apiKey}
	userPrompt := qroqMessage.Content

	response, err := groqClient.SendChatRequet("llama3-8b-8192", SystemPrompt, userPrompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
