//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("run", js.FuncOf(run))
	<-ch
}
