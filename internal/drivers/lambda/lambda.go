package lambda

import (
	"context"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/services"
	"github.com/exanubes/typedef/internal/usecase"
)

func Start(ctx context.Context, payload domain.CodegenRequest) (domain.GenerateCommandOutput, error) {
	input_service := services.NewInputService()
	codegen_service := configurator.New()
	usecase := usecase.NewGenerateUseCase(input_service, codegen_service)

	return usecase.Run(domain.GenerateCommandInput{
		Input:     payload.Data,
		Format:    payload.Format,
		InputType: payload.InputType,
	})
}
