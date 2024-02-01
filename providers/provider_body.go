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
	FunctionCall     ParameterConfig `json:"function_call"`
	MaxTokens        ParameterConfig `json:"max_tokens"`
	Temperature      ParameterConfig `json:"temperature"`
	TopP             ParameterConfig `json:"top_p"`
	N                ParameterConfig `json:"n"`
	Stream           ParameterConfig `json:"stream"`
	Stop             ParameterConfig `json:"stop"`
	PresencePenalty  ParameterConfig `json:"presence_penalty"`
	FrequencyPenalty ParameterConfig `json:"frequency_penalty"`
	LogitBias        ParameterConfig `json:"logit_bias"`
	User             ParameterConfig `json:"user"`
	Seed             ParameterConfig `json:"seed"`
	Tools            ParameterConfig `json:"tools"`
	ToolChoice       ParameterConfig `json:"tool_choice"`
	ResponseFormat   ParameterConfig `json:"response_format"`
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
	FinishReason string `json:"finish_reason"`
}

type ChatCompleteResponse struct {
	Id      string `json:"id"`
	Created int    `json:"created"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []ChatChoice `json:"choices"`
	*BaseResponse
}

type BaseResponse struct {
	Object string `json:"object"`
	Model  string `json:"model"`
}

type Params struct {
	Model            string             `json:"model"`
	Prompt           []string           `json:"prompt"`
	Messages         []Message          `json:"messages"`
	Functions        []Function         `json:"functions"`
	FunctionCall     interface{}        `json:"function_call"`
	MaxTokens        int                `json:"max_tokens"`
	Temperature      float64            `json:"temperature"`
	TopP             float64            `json:"top_p"`
	N                int                `json:"n"`
	Stream           bool               `json:"stream"`
	Logprobs         int                `json:"logprobs"`
	Echo             bool               `json:"echo"`
	Stop             []string           `json:"stop"`
	PresencePenalty  float64            `json:"presence_penalty"`
	FrequencyPenalty float64            `json:"frequency_penalty"`
	BestOf           int                `json:"best_of"`
	LogitBias        map[string]float64 `json:"logit_bias"`
	User             string             `json:"user"`
	Context          string             `json:"context"`
	Examples         []Example          `json:"examples"`
	TopK             int                `json:"top_k"`
	Tools            []Tool             `json:"tools"`
}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  JsonSchema `json:"parameters"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Message struct {
	// The role of the message sender
	Role string `json:"role"`
	// The content of the message
	Content []ContentType `json:"content"`
	// The name of the function to call, if any
	Name string `json:"name"`
	// The function call to make, if any
	FunctionCall interface{} `json:"function_call"`
	// The tool calls to make, if any
	ToolCalls interface{} `json:"tool_calls"`
	// Citation metadata
	CitationMetadata CitationMetadata `json:"citationMetadata"`
}

type JsonSchema map[string]interface{}

type CitationMetadata struct {
	CitationSources []CitationSource `json:"citationSources"`
}

type CitationSource struct {
	StartIndex *int   `json:"startIndex"`
	EndIndex   *int   `json:"endIndex"`
	URI        string `json:"uri"`
	License    string `json:"license"`
}

type ContentType struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageURL Image  `json:"image_url,omitempty"`
}

type Image struct {
	URL string `json:"url"`
}

type Example struct {
	Input  *Message `json:"input,omitempty"`
	Output *Message `json:"output,omitempty"`
}
