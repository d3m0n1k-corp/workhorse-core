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

func jsOf(value common.Response) js.Value {
	resp_json, err := json.Marshal(value)
	if err != nil {
		logrus.Error("chain_execute: failed to marshal response")
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}

func ChainExecute(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		err_str := "chain_execute: invalid number of arguments, expected 2"
		logrus.Error(err_str)
		return jsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}
	chainLinks := args[0].String()
	input := args[1].String()
	var request []chain.ConverterChainLink
	err := json.Unmarshal([]byte(chainLinks), &request)

	if err != nil {
		err_str := "chain_execute: invalid request format"
		logrus.Error(err_str)
		return jsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	logrus.Tracef("Executing chain with input %v and config %s", input, chainLinks)
	result, err := app.ExecuteChain(request, input)
	logrus.Tracef("Chain returned %v with error %v", result, err)

	var err_str string
	if err != nil {
		err_str = err.Error()
	}
	return jsOf(common.Response{
		Result: result,
		Error:  &err_str,
	})
}
