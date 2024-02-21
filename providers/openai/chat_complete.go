package openai

import "github.com/vidurkhanal/infuse/providers"

func NewOpenAIChatCompleteConfig() map[string]providers.ParameterConfig {
	return map[string]providers.ParameterConfig{
		"model": {
			Default:  "gpt-3.5-turbo",
			Param:    "model",
			Required: true,
		},
		"messages": {
			Param:   "messages",
			Default: " ",
		},
		"functions": {
			Param: "functions",
		},
		"function_call": {
			Param: "function_call",
		},
		"max_tokens": {
			Param:   "max_tokens",
			Default: 100,
			Min:     0,
		},
		"temperature": {
			Param:   "temperature",
			Default: 1,
			Min:     0,
			Max:     2,
		},
		"top_p": {
			Param:   "top_p",
			Default: 1,
			Min:     0,
			Max:     1,
		},
		"n": {
			Param:   "n",
			Default: 1,
		},
		"stream": {
			Param:   "stream",
			Default: false,
		},
		"stop": {
			Param: "stop",
		},
		"presence_penalty": {
			Param: "presence_penalty",
			Min:   -2,
			Max:   2,
		},
		"frequency_penalty": {
			Param: "frequency_penalty",
			Min:   -2,
			Max:   2,
		},
		"logit_bias": {
			Param: "logit_bias",
		},
		"user": {
			Param: "user",
		},
		"seed": {
			Param: "seed",
		},
		"tools": {
			Param: "tools",
		},
		"tool_choice": {
			Param: "tool_choice",
		},
		"response_format": {
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
