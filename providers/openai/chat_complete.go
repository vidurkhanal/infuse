package openai

import "github.com/vidurkhanal/infuse/providers"

type OpenAIChatCompleteConfig struct {
	Model            providers.ParameterConfig `json:"model"`
	Messages         providers.ParameterConfig `json:"messages"`
	Functions        providers.ParameterConfig `json:"functions"`
	FunctionCall     providers.ParameterConfig `json:"function_call"`
	MaxTokens        providers.ParameterConfig `json:"max_tokens"`
	Temperature      providers.ParameterConfig `json:"temperature"`
	TopP             providers.ParameterConfig `json:"top_p"`
	N                providers.ParameterConfig `json:"n"`
	Stream           providers.ParameterConfig `json:"stream"`
	Stop             providers.ParameterConfig `json:"stop"`
	PresencePenalty  providers.ParameterConfig `json:"presence_penalty"`
	FrequencyPenalty providers.ParameterConfig `json:"frequency_penalty"`
	LogitBias        providers.ParameterConfig `json:"logit_bias"`
	User             providers.ParameterConfig `json:"user"`
	Seed             providers.ParameterConfig `json:"seed"`
	Tools            providers.ParameterConfig `json:"tools"`
	ToolChoice       providers.ParameterConfig `json:"tool_choice"`
	ResponseFormat   providers.ParameterConfig `json:"response_format"`
}

func NewOpenAIChatCompleteConfig() *OpenAIChatCompleteConfig {
	return &OpenAIChatCompleteConfig{
		Model: providers.ParameterConfig{
			Default:  "gpt-3.5-turbo",
			Param:    "model",
			Required: true,
		},
		Messages: providers.ParameterConfig{
			Param:   "messages",
			Default: " ",
		},
		Functions: providers.ParameterConfig{
			Param: "functions",
		},
		FunctionCall: providers.ParameterConfig{
			Param: "function_call",
		},
		MaxTokens: providers.ParameterConfig{
			Param:   "max_tokens",
			Default: 100,
			Min:     0,
		},
		Temperature: providers.ParameterConfig{
			Param:   "temperature",
			Default: 1,
			Min:     0,
			Max:     2,
		},
		TopP: providers.ParameterConfig{
			Param:   "top_p",
			Default: 1,
			Min:     0,
			Max:     1,
		},
		N: providers.ParameterConfig{
			Param:   "n",
			Default: 1,
		},
		Stream: providers.ParameterConfig{
			Param:   "stream",
			Default: false,
		},
		Stop: providers.ParameterConfig{
			Param: "stop",
		},
		PresencePenalty: providers.ParameterConfig{
			Param: "presence_penalty",
			Min:   -2,
			Max:   2,
		},
		FrequencyPenalty: providers.ParameterConfig{
			Param: "frequency_penalty",
			Min:   -2,
			Max:   2,
		},
		LogitBias: providers.ParameterConfig{
			Param: "logit_bias",
		},
		User: providers.ParameterConfig{
			Param: "user",
		},
		Seed: providers.ParameterConfig{
			Param: "seed",
		},
		Tools: providers.ParameterConfig{
			Param: "tools",
		},
		ToolChoice: providers.ParameterConfig{
			Param: "tool_choice",
		},
		ResponseFormat: providers.ParameterConfig{
			Param: "response_format",
		},
	}
}
