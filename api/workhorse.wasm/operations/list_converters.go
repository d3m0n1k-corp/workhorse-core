//go:build js && wasm

package operations

import (
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/common"
	"workhorse-core/app"

	"github.com/sirupsen/logrus"
)

func ListConverters(this js.Value, args []js.Value) any {
	connList := app.ListConverters()
	logrus.Tracef("List of converters: %v", connList)
	return jsOf(common.Response{
		Result: connList,
		Error:  nil,
	})
}
