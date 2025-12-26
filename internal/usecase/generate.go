package usecase

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type GenerateUsecase struct {
	codegen domain.CodegenService
}

func NewGenerateUseCase(codegen domain.CodegenService) *GenerateUsecase {
	return &GenerateUsecase{
		codegen: codegen,
	}
}

func (usecase *GenerateUsecase) Run(cmd domain.GenerateCommandInput) (domain.GenerateCommandOutput, error) {
	response := domain.GenerateCommandOutput{}

	code, err := usecase.codegen.Execute(domain.CodegenOptions{
		Input:     cmd.Input,
		InputType: cmd.InputType,
		Format:    cmd.Format,
	})

	if err != nil {
		return response, fmt.Errorf("Failed execution of codegen service:\n| %w", err)
	}

	response.Format = cmd.Format
	response.Code = code

	return response, nil
}
