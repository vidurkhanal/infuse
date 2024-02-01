package openai

import "github.com/vidurkhanal/infuse/providers"

func NewOpenAIChatCompleteConfig() *providers.Params {
	return &providers.Params{
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

type OpenAIChatCompleteResponse struct {
	*providers.ChatCompleteResponse
	SystemFingerPrint string `json:"system_fingerprint"`
}

func OpenAIChatCompleteResponseTransform(response OpenAIChatCompleteResponse) OpenAIChatCompleteResponse {
	return response
}
