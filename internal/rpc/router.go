package rpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/usecase"
)

type RpcRouter struct {
	controllers map[string]ControllerHandler
}

func NewRouter() *RpcRouter {
	router := &RpcRouter{
		controllers: make(map[string]ControllerHandler),
	}
	router.register_methods()

	return router
}

func (router *RpcRouter) Get(method string) (ControllerHandler, bool) {
	handler, exists := router.controllers[method]
	return handler, exists
}

func (router *RpcRouter) register_methods() {
	router.controllers["codegen"] = func(id int, params json.RawMessage) (any, error) {
		var req codegenRPCRequest

		if err := json.Unmarshal(params, &req); err != nil {
			return nil, err
		}

		format, err := domain.ParseFormat(req.Format)

		if err != nil {
			return nil, err
		}
		input := strings.Trim(req.Input, " ")
		if input == "" {
			return nil, fmt.Errorf("Input cannot be empty")
		}

		codegen_service := configurator.New()
		usecase := usecase.NewGenerateUseCase(codegen_service)

		return usecase.Run(domain.GenerateCommandInput{
			Input:     input,
			InputType: req.InputType,
			Format:    format,
		})
	}
}
