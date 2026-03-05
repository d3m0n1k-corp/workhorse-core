package base

type BaseConfig interface {
	Validate() error
}
