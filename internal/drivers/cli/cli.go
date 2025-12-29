package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/infrastructure/clipboard"
	"github.com/exanubes/typedef/internal/infrastructure/readers"
	"github.com/exanubes/typedef/internal/infrastructure/targets"
	"github.com/exanubes/typedef/internal/services"
	"github.com/exanubes/typedef/internal/usecase"
)

var default_target = "cli"

func Start(ctx context.Context, args []string) error {
	cmd := flag.NewFlagSet("root", flag.ExitOnError)
	input_flag := cmd.String("input", "", "object structure that should be turned into a type definition or schema")
	target_flag := cmd.String("target", "clipboard", "delivery target for output e.g., cli, clipboard")
	output_path_flag := cmd.String("output-path", "", "path of the file where output should be saved")
	format_flag := cmd.String("format", "", "desired format for the output e.g., go, ts, ts-zod")

	cmd.Parse(args)
	clipboard_adapter := clipboard.New()

	if clipboard_adapter == nil {
		target_flag = &default_target
		fmt.Print("\nINFO: required dependencies for clipboard not detected. Falling back to cli target\n\n")
	}

	output_service := services.NewOutputService(targets.TargetFactory{}, clipboard_adapter)

	input_reader := readers.NewChainReader(
		readers.NewFlagReader(*input_flag),
		readers.NewStdinReader(),
		readers.NewClipboardReader(clipboard_adapter),
	)

	input_data, err := input_reader.Read()

	if err != nil {
		return err
	}

	format, err := domain.ParseFormat(*format_flag)

	if err != nil {
		return err
	}

	target, output_path := resolve_output(*target_flag, *output_path_flag)

	codegen_service := configurator.New()
	usecase := usecase.NewGenerateUseCase(codegen_service)

	output, err := usecase.Run(domain.GenerateCommandInput{
		Input:      input_data,
		Target:     target,
		OutputPath: output_path,
		Format:     format,
		InputType:  "json",
	})

	if err != nil {
		return fmt.Errorf("Failed execution of generate usecase:\n| %w", err)
	}

	err = output_service.Send(output.Code, domain.OutputOptions{Target: target, Path: output_path})

	if err != nil {
		return fmt.Errorf("Failed to send generated code to selected output '%s':\n| %w", target, err)
	}

	return nil
}

func resolve_output(target, output_path string) (string, string) {
	if output_path != "" {
		return "file", output_path
	}

	if target == "file" {
		if output_path == "" {
			return target, domain.DEFAULT_OUTPUT_PATH
		}

		return target, output_path
	}

	return target, ""
}
