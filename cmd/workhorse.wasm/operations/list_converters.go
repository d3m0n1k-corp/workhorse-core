//go:build js && wasm

package operations

import (
	"syscall/js"
	"workhorse-core/app"
	"workhorse-core/cmd/workhorse.wasm/common"

	"github.com/sirupsen/logrus"
)

func ListConverters(this js.Value, args []js.Value) any {
	connList := app.ListConverters()
	logrus.Tracef("List of converters: %v", connList)
	return common.JsOf(common.Response{
		Result: connList,
		Error:  nil,
	})
}
