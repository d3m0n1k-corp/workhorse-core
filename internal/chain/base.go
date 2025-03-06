package chain

type GenericChain interface {
	Execute() error
	Logs() []string
}
