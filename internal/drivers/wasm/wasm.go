package wasm

import "github.com/exanubes/typedef/internal/rpc"

func Start(input rpc.JSONRPCRequest) rpc.JSONRPCResponse {
	rpc_router := rpc.NewRouter()

	// TODO:
	// * marshall input string/byte slice into rpc.JSONRPCRequest
	// * get method handler from router
	// * run handler with rpc.JSONRPCRequest.Params
	// * create rpc.JSONRPCResponse
	// * return

	return rpc.JSONRPCResponse{}
}
