package basec

type BaseConverter interface {
	Apply() (any, error)
	Logs() []string
}
