//go:build js && wasm

package operations

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"

	"github.com/sirupsen/logrus"
)

func ExecuteConverter(this js.Value, args []js.Value) any {
	if len(args) != 3 {
		err_str := "execute_converter: invalid number of arguments, expected 3"
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	// Validate inputs
	name := args[0].String()
	if name == "" {
		err_str := "execute_converter: converter name cannot be empty"
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	input := args[1].String()
	// Note: input can be empty for some converters, so we don't validate it

	config := args[2].String()
	if config == "" {
		config = "{}"
	}

	// Validate config is valid JSON
	var configTest interface{}
	if err := json.Unmarshal([]byte(config), &configTest); err != nil {
		err_str := fmt.Sprintf("execute_converter: invalid JSON config: %v", err)
		logrus.Error(err_str)
		return common.JsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})
	}

	logrus.Tracef("Executing converter %s with input %v and config %s", name, input, config)
	result, err := app.ExecuteConverter(name, input, config)
	logrus.Tracef("Converter %s returned %s with error %v", name, result, err)

	var err_str string
	if err != nil {
		err_str = err.Error()
	}
	return common.JsOf(common.Response{
		Result: result,
		Error:  &err_str,
	})
}
