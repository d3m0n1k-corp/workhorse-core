//go:build js && wasm

package operations

import (
	"encoding/json"
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"

	"github.com/sirupsen/logrus"
)

func ExecuteConverter(this js.Value, args []js.Value) any {
	if len(args) != 3 {
		logrus.Error("execute_converter: invalid number of arguments")
		panic("execute_converter: invalid number of arguments")
	}
	name := args[0].String()
	input := args[1].String()
	config := args[2].String()

	logrus.Tracef("Executing converter %s with input %v and config %s", name, input, config)
	result, err := app.ExecuteConverter(name, input, config)
	logrus.Tracef("Converter %s returned %s with error %v", name, result, err)

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
