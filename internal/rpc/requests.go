package rpc

type codegenRPCRequest struct {
	Input     string `json:"input"`
	InputType string `json:"input_type"`
	Format    string `json:"format"`
}
