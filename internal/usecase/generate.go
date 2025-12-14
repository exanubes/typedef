package usecase

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type GenerateUsecase struct {
	codegen domain.CodegenService
	input   domain.InputResolver
}

func NewGenerateUseCase(input domain.InputResolver, codegen domain.CodegenService) *GenerateUsecase {
	return &GenerateUsecase{
		input:   input,
		codegen: codegen,
	}
}

func (usecase *GenerateUsecase) Run(cmd domain.GenerateCommandInput) (domain.GenerateCommandOutput, error) {
	response := domain.GenerateCommandOutput{}
	input, err := usecase.input.Resolve(cmd)

	if err != nil {
		return response, fmt.Errorf("Failed to resolve input:\n| %w", err)
	}

	code, err := usecase.codegen.Execute(domain.CodegenOptions{
		Input:     input.Data,
		InputType: cmd.InputType,
		Format:    input.Format,
	})

	if err != nil {
		return response, fmt.Errorf("Failed execution of codegen service:\n| %w", err)
	}

	response.Format = input.Format
	response.Code = code

	return response, nil
}
