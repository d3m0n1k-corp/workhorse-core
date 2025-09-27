//go:build js && wasm

package operations

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"
	"workhorse-core/internal/chain"

	"github.com/sirupsen/logrus"
)

func ChainExecute(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		err_str := "chain_execute: invalid number of arguments, expected 2"
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	chainLinks := args[0].String()
	if chainLinks == "" {
		err_str := "chain_execute: chain links cannot be empty"
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	input := args[1].String()
	// Note: input can be empty for some converters

	var request []chain.ConverterChainLink
	err := json.Unmarshal([]byte(chainLinks), &request)

	if err != nil {
		err_str := fmt.Sprintf("chain_execute: invalid JSON format for chain links: %v", err)
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	if len(request) == 0 {
		err_str := "chain_execute: chain links array cannot be empty"
		logrus.Error(err_str)
		return common.JsOf(common.Response{
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
	return common.JsOf(common.Response{
		Result: result,
		Error:  &err_str,
	})
}
