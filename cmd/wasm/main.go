//go:build js && wasm

package main

import "syscall/js"

func compile(this js.Value, args []js.Value) any {
	return nil
}

func main() {
	js.Global().Set("rpc", js.FuncOf(compile))
}
