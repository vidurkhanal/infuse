package constants

const POWERED_BY string = "infuse"

type Header struct {
	Mode     string
	Retries  string
	Provider string
	TraceID  string
	Cache    string
}

var HeaderKeys = Header{
	Mode:     "x-" + POWERED_BY + "-mode",
	Retries:  "x-" + POWERED_BY + "-retry-count",
	Provider: "x-" + POWERED_BY + "-provider",
	TraceID:  "x-" + POWERED_BY + "-trace-id",
	Cache:    "x-" + POWERED_BY + "-cache",
}

type Provider string

var (
	OPEN_AI       Provider = "openai"
	ANTHROPIC     Provider = "anthropic"
	DEEP_MIND     Provider = "deepmind"
	AZURE_OPEN_AI Provider = "azure-openai"
)

var Providers = []Provider{OPEN_AI, ANTHROPIC, DEEP_MIND}

const (
	StrategySingle      = "single"
	StrategyLoadBalance = "loadbalance"
	StrategyFallback    = "fallback"
	StrategyConcurrent  = "concurrent"

	CacheSimple   = "simple"
	CacheSemantic = "semantic"
)
