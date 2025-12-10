package cli

import (
	"context"
	"flag"

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

	generate_use_case := usecase.NewGenerateUseCase(
		services.NewOutputService(targets.TargetFactory{}),
		services.NewInputService(),
		configurator.New(),
	)

	return generate_use_case.Run(domain.GenerateCommand{
		Input:      *input_flag,
		Target:     *target_flag,
		OutputPath: *output_path_flag,
		Format:     *format_flag,
	})
}
