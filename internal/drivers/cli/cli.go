package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/domain"
)

func Start(ctx context.Context, args []string) error {
	cmd := flag.NewFlagSet("root", flag.ExitOnError)
	input_flag := cmd.String("input", "", "object structure that should be turned into a type definition or schema")
	// input_type_flag := cmd.String("type", "json", "type of the object structure")
	// target_flag := cmd.String("target", "cli", "delivery target for ouput e.g., cli, clipboard")
	// format_flag := cmd.String("format", "", "desired format for the output e.g., go, ts, ts-zod")

	cmd.Parse(args)

	input, err := parse_input(*input_flag)

	if err != nil {
		return err
	}

	// if *format_flag == "" {
	// 	return fmt.Errorf("--format flag is required")
	// }
	codegen_service, output_target := configurator.New()

	code, err := codegen_service.Execute(domain.CodegenOptions{
		Input:      input,
		InputType:  "json",
		OutputType: "golang",
	})

	if err != nil {
		return err
	}

	err = output_target.Send(code)

	return err
}

func parse_input(input string) (string, error) {
	if input != "" {
		return input, nil
	}
	stat, _ := os.Stdin.Stat()
	has_piped_input := (stat.Mode() & os.ModeCharDevice) == 0
	if has_piped_input {
		data, err := io.ReadAll(os.Stdin)
		return string(data), err
	}

	return "", fmt.Errorf("No input provided")
}
