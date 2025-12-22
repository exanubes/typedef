//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/exanubes/typedef/internal/drivers/wasm"
	"github.com/exanubes/typedef/internal/rpc"
)

var FALLBACK_ERROR_RESPONSE = `{"jsonrpc": "2.0","error":{"code": -32603, "message": "Internal error"}, "id": null}`

func compile(this js.Value, args []js.Value) any {
	raw_rpc_request_string := args[0].String()
	var request rpc.JSONRPCRequest

	if err := json.Unmarshal([]byte(raw_rpc_request_string), &request); err != nil {
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

func main() {
	js.Global().Set("rpc", js.FuncOf(compile))
}
