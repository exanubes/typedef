package usecase

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type GenerateUsecase struct {
	codegen domain.CodegenService
	output  domain.OutputService
	input   domain.InputResolver
}

func NewGenerateUseCase(output domain.OutputService, input domain.InputResolver, codegen domain.CodegenService) *GenerateUsecase {
	return &GenerateUsecase{
		output:  output,
		input:   input,
		codegen: codegen,
	}
}

func (usecase *GenerateUsecase) Run(cmd domain.GenerateCommand) error {
	input, err := usecase.input.Resolve(cmd)

	if err != nil {
		return fmt.Errorf("[GenerateUseCase] Failed to resolve input:\n| %w", err)
	}

	code, err := usecase.codegen.Execute(domain.CodegenOptions{
		Input:      input.Data,
		InputType:  "json",
		OutputType: input.Output.Target,
		Format:     input.Format,
	})

	if err != nil {
		return fmt.Errorf("[GenerateUseCase] Failed execution of codegen service:\n| %w", err)
	}

	err = usecase.output.Send(code, input.Output)

	if err != nil {
		return fmt.Errorf("[GenerateUseCase] Failed to send generated code to selected output '%s':\n| %w", input.Output.Target, err)
	}

	return nil
}
