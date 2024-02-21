package providers

type ProviderAPIConfig interface {
	Headers(string) map[string]string
	GetBaseURL() string
	GetEndpoint(string) (string, error)
}

type ParameterConfig struct {
	Param    string      `json:"param"`
	Default  interface{} `json:"default"`
	Min      int         `json:"min"`
	Max      int         `json:"max"`
	Required bool        `json:"required"`
}

type ChatChoice struct {
	Index        int    `json:"index"`
	Message      string `json:"message"`
	FinishReason string `json:"finishReason"`
}

type ChatCompleteResponse struct {
	Id      string `json:"id"`
	Created int    `json:"created"`
	Usage   struct {
		PromptTokens     int `json:"promptTokens"`
		CompletionTokens int `json:"completionTokens"`
		TotalTokens      int `json:"totalTokens"`
	} `json:"usage"`
	Choices []ChatChoice `json:"choices"`
	*BaseResponse
}

type BaseResponse struct {
	Object string `json:"object"`
	Model  string `json:"model"`
}
