package openai

import "github.com/vidurkhanal/infuse/providers"

type OpenAIConfig struct {
	ApiConfig          providers.ProviderAPIConfig
	ChatCompleteConfig *OpenAIChatCompleteConfig
}

func NewOpenAIConfig() *OpenAIConfig {
	return &OpenAIConfig{
		ApiConfig:          NewOpenAIAPIConfig(),
		ChatCompleteConfig: NewOpenAIChatCompleteConfig(),
	}
}
