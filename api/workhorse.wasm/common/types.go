//go:build js && wasm

package common

import (
	"encoding/json"
	"syscall/js"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Error  *string
	Result any
}

// JsOf marshals a Response to a JavaScript value
func JsOf(value Response) js.Value {
	resp_json, err := json.Marshal(value)
	if err != nil {
		logrus.Error("failed to marshal response to JSON")
		panic(err)
	}
	return js.ValueOf(string(resp_json))
}
