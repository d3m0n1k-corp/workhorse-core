//go:build js && wasm

package main

type Response struct {
	Error  *string
	Result any
}
