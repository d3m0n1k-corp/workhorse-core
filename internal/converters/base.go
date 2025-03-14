package converters

type BaseConverter interface {
	Apply(input any) (any, error)
	InputType() string
	OutputType() string
}

type BaseConfig interface {
	Validate(in map[string]any) error
}
