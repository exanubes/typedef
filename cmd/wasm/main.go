//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/exanubes/typedef/internal/drivers/wasm"
	"github.com/exanubes/typedef/internal/rpc"
)

var FALLBACK_ERROR_RESPONSE = `{"jsonrpc": "2.0","error":{"code": -32603, "message": "Internal error"}, "id": null}`
var inbox = make(chan js.Value, 64)

// NOTE: declaring in package scope to prevent GC from clearing them
var handler_func js.Func
var post_message js.Value

func compile(raw_request string) any {
	var request rpc.JSONRPCRequest

	if err := json.Unmarshal([]byte(raw_request), &request); err != nil {
		response := rpc.JSONRPCResponse{
			Version: "2.0",
			ID:      0,
			Error: &rpc.RPCError{
				Code:    rpc.ParseError,
				Message: err.Error(),
			},
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			return FALLBACK_ERROR_RESPONSE
		}

		return string(bytes)
	}

	response := wasm.Start(request)

	bytes, err := json.Marshal(response)
	if err != nil {
		return FALLBACK_ERROR_RESPONSE
	}

	return string(bytes)
}

func handler(this js.Value, args []js.Value) any {
	inbox <- args[0]
	return nil
}

func main() {
	handler_func = js.FuncOf(handler)
	post_message = js.Global().Get("postMessage")
	js.Global().Set("rpc", handler_func)

	go func() {
		for msg := range inbox {
			result := compile(msg.String())
			post_message.Invoke(js.ValueOf(result))
		}
	}()

	// NOTE: blocking main func from returning and closing process
	select {}
}
