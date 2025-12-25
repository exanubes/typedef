package rpc

import (
	"encoding/json"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/services"
	"github.com/exanubes/typedef/internal/usecase"
)

type RpcRouter struct {
	methods map[string]MethodHandler
}

func NewRouter() *RpcRouter {
	router := &RpcRouter{
		methods: make(map[string]MethodHandler),
	}
	router.register_methods()

	return router
}

func (router *RpcRouter) Get(method string) (MethodHandler, bool) {
	handler, exists := router.methods[method]
	return handler, exists
}

func (router *RpcRouter) register_methods() {
	router.methods["codegen"] = func(id int, params json.RawMessage) (any, error) {
		var cmd domain.GenerateCommandInput

		if err := json.Unmarshal(params, &cmd); err != nil {
			return nil, err
		}

		input_service := services.NewInputService()
		codegen_service := configurator.New()
		usecase := usecase.NewGenerateUseCase(input_service, codegen_service)

		return usecase.Run(cmd)
	}
}
