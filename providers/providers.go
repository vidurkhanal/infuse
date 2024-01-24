package providers

type ParameterConfig struct {
	Param    string      `json:"param"`
	Default  interface{} `json:"default"`
	Min      int         `json:"min"`
	Max      int         `json:"max"`
	Required bool        `json:"required"`
}
