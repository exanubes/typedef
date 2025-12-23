package wasm

import "github.com/exanubes/typedef/internal/rpc"

func Start(input rpc.JSONRPCRequest) rpc.JSONRPCResponse {
	rpc_router := rpc.NewRouter()

	handler, exists := rpc_router.Get(input.Method)

	if !exists {
		return rpc.JSONRPCResponse{
			Version: "2.0",
			ID:      input.ID,
			Error: &rpc.RPCError{
				Code:    rpc.MethodNotFound,
				Message: "Method not found",
			},
		}
	}

	response, err := handler(input.ID, input.Params)

	if err != nil {
		return rpc.JSONRPCResponse{
			Version: "2.0",
			ID:      input.ID,
			Error: &rpc.RPCError{
				Code:    rpc.InvalidRequest,
				Message: err.Error(),
			},
		}
	}

	return rpc.JSONRPCResponse{
		Version: "2.0",
		ID:      input.ID,
		Result:  response,
	}
}
