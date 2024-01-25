package openai

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

func (c *OpenAIAPIConfig) GetEndpoint() string {
	return c.BaseURL + c.ChatComplete
}

func NewOpenAIAPIConfig() *OpenAIAPIConfig {
	return &OpenAIAPIConfig{
		BaseURL:      "https://api.openai.com/v1",
		ChatComplete: "/chat/completions",
		Embed:        "/embeddings",
	}
}
