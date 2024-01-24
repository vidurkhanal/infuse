package openai

type OpenAIConfig struct {
	ApiConfig *OpenAIAPIConfig
}

var OpenAI = &OpenAIConfig{ApiConfig: NewOpenAIAPIConfig()}
