package rpc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Server struct {
	reader  *bufio.Scanner
	writer  *json.Encoder
	mutex   sync.Mutex
	methods Router
}

func NewServer(router Router) *Server {
	return &Server{
		reader:  bufio.NewScanner(os.Stdin),
		writer:  json.NewEncoder(os.Stdout),
		methods: router,
	}
}

func (server *Server) Start() error {
	for server.reader.Scan() {
		var request JSONRPCRequest

		if err := json.Unmarshal(server.reader.Bytes(), &request); err != nil {
			server.send_error(request.ID, ParseError, fmt.Sprintf("[RPC Server] %s", err.Error()))
			continue
		}

		go server.handle_request(request)
	}

	return server.reader.Err()
}

func (server *Server) send_error(id int, err_code int, message string) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	response := JSONRPCResponse{
		Version: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    err_code,
			Message: message,
		},
	}

	server.writer.Encode(response)
}

func (server *Server) handle_request(request JSONRPCRequest) {
	handler, exists := server.methods.Get(request.Method)

	if !exists {
		server.send_error(request.ID, MethodNotFound, fmt.Sprintf("method '%s' not found", request.Method))
		return
	}

	result, err := handler(request.ID, request.Params)
	if err != nil {
		server.send_error(request.ID, InternalError, fmt.Sprintf("Unhandled Exception: %s", err.Error()))
		return
	}

	server.send_response(request.ID, result)
}

func (server *Server) send_response(id int, result any) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	response := JSONRPCResponse{
		ID:      id,
		Version: "2.0",
		Result:  result,
	}

	server.writer.Encode(response)
}
