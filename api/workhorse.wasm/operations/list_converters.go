//go:build js && wasm

package operations

import (
	"encoding/json"
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"
)

func List_connectors(this js.Value, args []js.Value) any {
	connList := app.ListConverters()
	// l.GlobalStream.Logf("Found %d connectors", len(connList))
	response_object := common.Response{
		Result: connList,
		Error:  nil,
	}
	resp_json, err := json.Marshal(response_object)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}
