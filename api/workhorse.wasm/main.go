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
	logrus.SetLevel(logrus.InfoLevel)
}

func registerFunctions() {
	js.Global().Set("list_connectors", js.FuncOf(operations.List_connectors))
	js.Global().Set("execute_converter", js.FuncOf(operations.Execute_converter))
	js.Global().Get("console").Call("log", "WASM Initialized and Ready")
}

func waitTillDone() {
	<-make(chan bool)
}
