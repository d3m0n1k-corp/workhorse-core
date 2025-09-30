//go:build js && wasm

package logging

import "syscall/js"

type JsLogWriter struct {
}

func (w JsLogWriter) Write(p []byte) (n int, err error) {
	js.Global().Get("console").Call("log", string(p))
	return len(p), nil
}
