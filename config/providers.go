package config

import (
	"sync"

	"github.com/vidurkhanal/infuse/providers/openai"
)

type ProvidersConfig struct {
	OpenAIConfig *openai.OpenAIConfig
}

var (
	instance *ProvidersConfig
	once     sync.Once
)

func Providers() *ProvidersConfig {
	once.Do(func() {
		instance = &ProvidersConfig{
			OpenAIConfig: openai.NewOpenAIConfig(),
		}
	})
	return instance
}
