//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"
	"workhorse-core/app"
)

func list_connectors(this js.Value, args []js.Value) any {
	response_object := Response{
		Result: app.ListConvertersInJSON(),
		Error:  nil,
	}
	resp_json, err := json.Marshal(response_object)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}
