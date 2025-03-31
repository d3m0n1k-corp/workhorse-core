//go:build js && wasm

package operations

import (
	"encoding/json"
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"
	"workhorse-core/internal/chain"

	"github.com/sirupsen/logrus"
)

func ChainExecute(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		logrus.Error("chain_execute: invalid number of arguments")
		panic("chain_execute: invalid number of arguments")
	}
	chainLinks := args[0].String()
	input := args[1].String()
	var request []chain.ConverterChainLink
	err := json.Unmarshal([]byte(chainLinks), &request)

	if err != nil {
		logrus.Error("chain_execute: invalid request format")
		panic("chain_execute: invalid request format")
	}

	logrus.Tracef("Executing chain with input %v and config %s", input, chainLinks)
	result, err := app.ExecuteChain(request, input)
	logrus.Tracef("Chain returned %v with error %v", result, err)

	var err_str string
	if err != nil {
		err_str = err.Error()
	}
	response_object := common.Response{
		Result: result,
		Error:  &err_str,
	}

	resp_json, err := json.Marshal(response_object)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}
