package app

import (
	"fmt"
	"reflect"
	"workhorse-core/internal/chain"

	"github.com/sirupsen/logrus"
)

type ChainResponse struct {
	Output any    `json:"output"`
	Error  string `json:"error"`
}

func ExecuteChain(chainLinks []chain.ConverterChainLink, input string) ([]*ChainResponse, error) {
	cl, err := chain.NewConverterListFromJSON(chainLinks)
	if err != nil {
		logrus.Errorf("Error creating converter list: %v", err)
		return nil, err
	}
	logrus.Trace("Converter list created successfully")
	err = cl.Validate()
	if err != nil {
		logrus.Errorf("Error validating converter list: %v", err)
		return nil, err
	}
	logrus.Trace("Converter list validated successfully")

	head := cl.Head()
	response := make([]*ChainResponse, cl.Length())
	chain_in := input
	dbg_i := 0
	for head != nil {
		logrus.Tracef("Executing converter %d: %s", dbg_i, reflect.TypeOf(head.Value))
		out, err := head.Value.Apply(chain_in)
		resp := ChainResponse{
			Output: out,
			Error:  fmt.Sprintf("%v", err),
		}

		response = append(response, &resp)

		if err != nil {
			break
		}
		chain_in = out.(string)
		head = head.Next
		dbg_i++
	}

	return response, err
}
