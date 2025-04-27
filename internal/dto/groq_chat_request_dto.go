package dto

type GroqChatRequest struct {
	Messages    []GroqMessage `json:"messages"`
	Model       string        `json:"model"`
	Temperature int           `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
	TopP        int           `json:"top_p"`
	Stream      bool          `json:"stream"`
	Stop        interface{}   `json:"stop"`
}
