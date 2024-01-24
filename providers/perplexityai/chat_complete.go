package perplexityai

type PerplexityAIChoice struct {
	Message struct {
		Role    string
		Content string
	}
	Delta struct {
		Role    string
		Content string
	}
	Index        int
	FinishReason string
}

type PerplexityAIResponse struct {
	ID      string
	Model   string
	Object  string
	Created int64
	Choices []PerplexityAIChoice
	Usage   struct {
		PromptTokens     int
		CompletionTokens int
		TotalTokens      int
	}
}

type PerplexityAIError struct {
	Error struct {
		Message string
		Type    string
		Code    int
	}
}
