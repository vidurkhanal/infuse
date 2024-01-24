package providers

type ParameterConfig struct {
	Param     string
	Default   interface{}
	Min       *int
	Max       *int
	Required  bool
	Transform func(interface{}) interface{}
}
