package base

type BaseConverter interface {
	Apply(input any) (any, error)
	InputType() string
	OutputType() string
}
