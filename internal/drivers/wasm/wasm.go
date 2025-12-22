package wasm

import "github.com/exanubes/typedef/internal/drivers/rpc"

func Start(input []byte) {
	rpc_router := rpc.NewRouter()

	// TODO:
	// * marshall input string/byte slice into rpc.JSONRPCRequest
	// * get method handler from router
	// * run handler with rpc.JSONRPCRequest.Params
	// * create rpc.JSONRPCResponse
	// * return
}
