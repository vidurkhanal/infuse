package perplexityai

type PerplexityAIConfig struct {
	BaseURL      string
	ChatComplete string
}

func (c *PerplexityAIConfig) Headers(API_KEY string) map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + API_KEY,
	}
}

func NewPerplexityAIConfig() *PerplexityAIConfig {
	return &PerplexityAIConfig{
		BaseURL:      "https://api.perplexity.ai",
		ChatComplete: "/v1/chat-complete",
	}
}
