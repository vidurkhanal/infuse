package providers

type ProviderAPIConfig interface {
	Headers(string) map[string]string
	GetBaseURL() string
	GetEndpoint() string
}

type ProviderFeatConfig struct {
	Model            ParameterConfig `json:"model"`
	Messages         ParameterConfig `json:"messages"`
	Functions        ParameterConfig `json:"functions"`
	FunctionCall     ParameterConfig `json:"functionCall"`
	MaxTokens        ParameterConfig `json:"maxTokens"`
	Temperature      ParameterConfig `json:"temperature"`
	TopP             ParameterConfig `json:"topP"`
	N                ParameterConfig `json:"n"`
	Stream           ParameterConfig `json:"stream"`
	Stop             ParameterConfig `json:"stop"`
	PresencePenalty  ParameterConfig `json:"presencePenalty"`
	FrequencyPenalty ParameterConfig `json:"frequencyPenalty"`
	LogitBias        ParameterConfig `json:"logitBias"`
	User             ParameterConfig `json:"user"`
	Seed             ParameterConfig `json:"seed"`
	Tools            ParameterConfig `json:"tools"`
	ToolChoice       ParameterConfig `json:"toolChoice"`
	ResponseFormat   ParameterConfig `json:"responseFormat"`
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
