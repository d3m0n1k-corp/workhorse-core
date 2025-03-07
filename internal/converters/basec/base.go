package basec

type BaseConverter interface {
	Apply() (any, error)
	Logs() []string
	Inputtype() *interface{}
	Outputtype() *interface{}
}
