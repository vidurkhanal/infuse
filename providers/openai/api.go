package openai

import "errors"

type OpenAIAPIConfig struct {
	BaseURL      string
	ChatComplete string
	Embed        string
}

func (c *OpenAIAPIConfig) Headers(API_KEY string) map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + API_KEY,
	}
}

func (c *OpenAIAPIConfig) GetBaseURL() string {
	return c.BaseURL
}

func (c *OpenAIAPIConfig) GetEndpoint(fn string) (string, error) {
	switch fn {
	case "chat":
		return c.BaseURL + c.ChatComplete, nil
	case "embed":
		return c.BaseURL + c.Embed, nil
	default:
		return "", errors.New("Invalid function")
	}
}

func NewOpenAIAPIConfig() *OpenAIAPIConfig {
	return &OpenAIAPIConfig{
		BaseURL:      "https://api.openai.com/v1",
		ChatComplete: "/chat/completions",
		Embed:        "/embeddings",
	}
}
