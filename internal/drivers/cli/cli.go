package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/exanubes/typedef/internal/app/configurator"
	"github.com/exanubes/typedef/internal/app/generator"
	"github.com/exanubes/typedef/internal/domain"
)

func Start(ctx context.Context, args []string) error {
	cmd := flag.NewFlagSet("root", flag.ExitOnError)
	input_flag := cmd.String("input", "", "object structure that should be turned into a type definition or schema")
	// input_type_flag := cmd.String("type", "json", "type of the object structure")
	target_flag := cmd.String("target", "cli", "delivery target for ouput e.g., cli, clipboard")
	output_path_flag := cmd.String("output-path", "", "path of the file where output should be saved")
	format_flag := cmd.String("format", "", "desired format for the output e.g., go, ts, ts-zod")

	cmd.Parse(args)

	input, err := parse_input(*input_flag)

	if err != nil {
		return err
	}

	format, err := parse_format(*format_flag)

	if err != nil {
		return err
	}

	if format == -1 {
		return fmt.Errorf("--format flag is required")
	}

	target, output_path := parse_target(*target_flag, *output_path_flag)

	codegen_service, output_target := configurator.New(configurator.Options{
		OutputTarget:         target,
		OutputTargetFilepath: output_path,
	})

	code, err := codegen_service.Execute(domain.CodegenOptions{
		Input:      input,
		InputType:  "json",
		OutputType: "golang",
		Format:     int(format),
	})

	if err != nil {
		return err
	}

	err = output_target.Send(code)

	return err
}

func parse_target(target string, output_path string) (string, string) {
	if output_path != "" {
		return "file", output_path
	}

	if target == "file" {
		if output_path == "" {
			return target, "./typedef.txt"
		}

		return target, output_path
	}

	return target, ""
}

func parse_format(input string) (generator.Format, error) {
	switch input {
	case "go", "golang":
		return generator.GOLANG, nil

	case "typescript", "ts":
		return generator.TYPESCRIPT, nil

	case "zod", "ts-zod", "ts_zod", "typescript_zod", "typescript-zod":
		return generator.ZOD, nil
	case "jsdoc":
		return generator.JSDOC, nil
	}

	return -1, fmt.Errorf("%s is not a valid format", input)
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
