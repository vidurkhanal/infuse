package config

import "github.com/vidurkhanal/infuse/constants"

type Config struct {
	/** maybe "single" | "fallback" | "loadbalance" | "scientist" */
	Mode     string        `json:"mode"`
	Targets  []Target      `json:"targets"`
	Cache    CacheSettings `json:"cache"`
	Retry    RetrySettings `json:"retry"`
	Strategy Strategy      `json:"strategy"`
}

type Target struct {
	Strategy Strategy `json:"strategy"`
	/** The name of the provider. */
	Provider constants.Provider `json:"provider"`
	/** The name of the API key for the provider. */
	VirtualKey string `json:"virtualKey"`
	/** The API key for the provider. */
	ApiKey string `json:"apiKey"`
	/** The weight of the provider, used for load balancing. */
	Weight int `json:"weight"`
	/** The retry settings for the provider. */
	Retry RetrySettings `json:"retry"`
	/** Azure specific */
	ResourceName string `json:"resourceName"`
	DeploymentId string `json:"deploymentId"`
	ApiVersion   string `json:"apiVersion"`
	AdAuth       string `json:"adAuth"`
	/** provider option index picked based on weight in loadbalance mode */
	Index   int           `json:"index"`
	Cache   CacheSettings `json:"cache"`
	Targets []Target      `json:"targets"`
}

type RetrySettings struct {
	/** The maximum number of retry attempts. */
	Attempts int `json:"attempts"`
	/** The HTTP status codes on which to retry. */
	OnStatusCodes []int `json:"onStatusCodes"`
}

type CacheSettings struct {
	Mode   string `json:"mode"` // "none" | "simple" | "redis"
	MaxAge int    `json:"maxAge"`
}

type Strategy struct {
	Mode          string `json:"mode"`
	OnStatusCodes []int  `json:"onStatusCodes"`
}

type Params struct {
	Model            string             `json:"model,omitempty"`
	Prompt           []string           `json:"prompt,omitempty"`
	Messages         []Message          `json:"messages,omitempty"`
	Functions        []Function         `json:"functions,omitempty"`
	FunctionCall     string             `json:"function_call,omitempty"`
	MaxTokens        int                `json:"max_tokens,omitempty"`
	Temperature      float64            `json:"temperature,omitempty"`
	TopP             float64            `json:"top_p,omitempty"`
	N                int                `json:"n,omitempty"`
	Stream           bool               `json:"stream,omitempty"`
	Logprobs         int                `json:"logprobs,omitempty"`
	Echo             bool               `json:"echo,omitempty"`
	PresencePenalty  float64            `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64            `json:"frequency_penalty,omitempty"`
	BestOf           int                `json:"best_of,omitempty"`
	LogitBias        map[string]float64 `json:"logit_bias,omitempty"`
	User             string             `json:"user,omitempty"`
	Context          string             `json:"context,omitempty"`
	Examples         []Example          `json:"examples,omitempty"`
	TopK             int                `json:"top_k,omitempty"`
	Tools            []Tool             `json:"tools,omitempty"`
}

type Message struct {
	Role             string            `json:"role"` // system, user, assistant, function
	Content          []ContentType     `json:"content,omitempty"`
	Name             string            `json:"name,omitempty"`
	FunctionCall     interface{}       `json:"function_call,omitempty"`
	ToolCalls        interface{}       `json:"tool_calls,omitempty"`
	CitationMetadata *CitationMetadata `json:"citationMetadata,omitempty"`
}

type ContentType struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageURL Image  `json:"image_url,omitempty"`
}

type Image struct {
	URL string `json:"url"`
}

type CitationMetadata struct {
	CitationSources []CitationSource `json:"citationSources,omitempty"`
}

type CitationSource struct {
	StartIndex int    `json:"startIndex,omitempty"`
	EndIndex   int    `json:"endIndex,omitempty"`
	URI        string `json:"uri,omitempty"`
	License    string `json:"license,omitempty"`
}

type JsonSchema map[string]interface{}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Parameters  JsonSchema `json:"parameters,omitempty"`
}

type Tool struct {
	Type     string    `json:"type"`
	Function *Function `json:"function,omitempty"`
}

type Example struct {
	Input  Message `json:"input"`
	Output Message `json:"output"`
}

/**
 * The full structure of the request body.
 * @interface
 */
type FullRequestBody struct {
	/** The configuration for handling the request. */
	Config Config `json:"config"`
	/** The parameters for the request. */
	Params Params `json:"params"`
}
