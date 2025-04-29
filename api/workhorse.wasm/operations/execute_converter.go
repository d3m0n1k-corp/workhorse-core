//go:build js && wasm

package operations

import (
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"

	"github.com/sirupsen/logrus"
)

func ExecuteConverter(this js.Value, args []js.Value) any {
	if len(args) != 3 {
		err_str := "execute_converter: invalid number of arguments"
		logrus.Error(err_str)
		return jsOf(common.Response{
			Result: nil,
			Error:  &err_str,
		})

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
	return jsOf(common.Response{
		Result: result,
		Error:  &err_str,
	})
}
