//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"
	"workhorse-core/app"
)

func list_connectors(this js.Value, args []js.Value) any {
	response_object := Response{
		Result: app.ListConverters(),
		Error:  nil,
	}
	resp_json, err := json.Marshal(response_object)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}

func execute_converter(this js.Value, args []js.Value) any {
	if len(args) != 3 {
		panic("execute_converter: invalid number of arguments")
	}
	name := args[0].String()
	input := args[1].String()
	config := args[2].String()

	result, err := app.ExecuteConverter(name, input, config)
	err_str := err.Error()
	response_object := Response{
		Result: result,
		Error:  &err_str,
	}
	resp_json, err := json.Marshal(response_object)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}
