package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/infrastructure/targets"
	"github.com/exanubes/typedef/internal/services"
	"github.com/exanubes/typedef/internal/usecase"
)

func Start(ctx context.Context, args []string) error {
	cmd := flag.NewFlagSet("root", flag.ExitOnError)
	input_flag := cmd.String("input", "", "object structure that should be turned into a type definition or schema")
	target_flag := cmd.String("target", "clipboard", "delivery target for ouput e.g., cli, clipboard")
	output_path_flag := cmd.String("output-path", "", "path of the file where output should be saved")
	format_flag := cmd.String("format", "", "desired format for the output e.g., go, ts, ts-zod")
	cmd.Parse(args)

	output_service := services.NewOutputService(targets.TargetFactory{})

	input_service := services.NewInputService()
	codegen_service := configurator.New()
	usecase := usecase.NewGenerateUseCase(input_service, codegen_service)

	output, err := usecase.Run(domain.GenerateCommandInput{
		Input:      *input_flag,
		Target:     *target_flag,
		OutputPath: *output_path_flag,
		Format:     *format_flag,
	})

	if err != nil {
		return fmt.Errorf("Failed execution of generate usecase:\n| %w", err)
	}

	target, path := input_service.ResolveTarget(*target_flag, *output_path_flag)
	err = output_service.Send(output.Code, domain.OutputOptions{Target: target, Path: path})

	if err != nil {
		return fmt.Errorf("Failed to send generated code to selected output '%s':\n| %w", target, err)
	}

	return nil
}
