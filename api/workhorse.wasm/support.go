//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func errorOut(errors []error) js.Value {
	vs := make([]any, 0, len(errors))

	for _, r := range errors {
		vs = append(vs, string(r.Error()))
	}
	return js.ValueOf(vs)
}

func run(this js.Value, args []js.Value) any {
	errors := []error{}
	var t any
	fmt.Printf("args: %v\n", args)
	fmt.Printf("this: %v\n", this)

	err := json.Unmarshal([]byte(args[0].String()), &t)

	if err != nil {
		return errorOut(append(errors, fmt.Errorf("Null JSON")))
	}

	_, ok := t.(map[string]any)

	if !ok {
		return errorOut(append(errors, fmt.Errorf("Invalid JSON")))
	}

	return errorOut(errors)
}
