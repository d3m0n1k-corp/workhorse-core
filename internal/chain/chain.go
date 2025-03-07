package chain

import (
	"workhorse-core/internal/converters"
	"workhorse-core/internal/converters/basec"
)

type Chain struct {
	in     any
	out    any
	logs   []string
	Blocks []*basec.BaseConverter
}

func (c *Chain) Logs() []string {
	return c.logs
}

func NewChain(in any, chain_blocks []string) (GenericChain, error) {

	blocks := []basec.BaseConverter{}
	for _, block := range chain_blocks {
		conv, err := converters.GetConverter(block)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, conv)
	}

	chain := &Chain{
		in:     in,
		out:    nil,
		logs:   []string{},
		Blocks: []*basec.BaseConverter{},
	}
	return chain, nil
}

func (c *Chain) Execute() error {
	panic("implement me")
}
