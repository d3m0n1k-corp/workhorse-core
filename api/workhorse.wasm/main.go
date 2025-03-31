//go:build js && wasm

package main

import (
	"syscall/js"
	"workhorse-core/api/workhorse.wasm/logging"
	"workhorse-core/api/workhorse.wasm/operations"

	"github.com/sirupsen/logrus"
)

func main() {
	registerLogger()
	registerFunctions()
	waitTillDone()
}

func registerLogger() {
	writer := logging.JsLogWriter{}
	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.TraceLevel)
}

func registerFunctions() {
	js.Global().Set("list_converters", js.FuncOf(operations.ListConverters))
	js.Global().Set("execute_converter", js.FuncOf(operations.ExecuteConverter))
	js.Global().Set("chain_execute", js.FuncOf(operations.ChainExecute))
	js.Global().Get("console").Call("log", "WASM Initialized and Ready")
}

func waitTillDone() {
	<-make(chan bool)
}
