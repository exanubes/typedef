package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/exanubes/typedef/internal/domain"
	"golang.design/x/clipboard"
)

type InputService struct{}

func NewInputService() InputService {
	return InputService{}
}

// TODO: make configurable
var default_path_target = "./typedef.txt"

func (service InputService) Resolve(input domain.GenerateCommandInput) (domain.ResolvedInput, error) {
	data, err := service.parse_input(input.Input)
	if err != nil {
		return domain.ResolvedInput{}, err
	}

	format, err := service.parse_format(input.Format)

	if err != nil {
		return domain.ResolvedInput{}, err
	}

	if format == -1 {
		return domain.ResolvedInput{}, fmt.Errorf("--format flag is required")
	}

	target, output_path := service.parse_target(input.Target, input.OutputPath)

	return domain.ResolvedInput{
		Data:   data,
		Format: format,
		Output: domain.OutputOptions{
			Target: target,
			Path:   output_path,
		},
	}, nil
}

func (service InputService) ResolveTarget(target, path string) (string, string) {
	return service.parse_target(target, path)
}

func (service InputService) parse_format(input string) (domain.Format, error) {
	switch input {
	case "go", "golang":
		return domain.GOLANG, nil

	case "typescript", "ts":
		return domain.TYPESCRIPT, nil

	case "zod", "ts-zod", "ts_zod", "typescript_zod", "typescript-zod":
		return domain.ZOD, nil
	case "jsdoc":
		return domain.JSDOC, nil
	}

	return -1, fmt.Errorf("%s is not a valid format", input)
}

func (service InputService) parse_target(target string, output_path string) (string, string) {
	if output_path != "" {
		return "file", output_path
	}

	if target == "file" {
		if output_path == "" {
			return target, default_path_target
		}

		return target, output_path
	}

	return target, ""
}

func (service InputService) parse_input(input string) (string, error) {
	if input != "" {
		return input, nil
	}

	stat, _ := os.Stdin.Stat()
	has_piped_input := (stat.Mode() & os.ModeCharDevice) == 0
	if has_piped_input {
		data, err := io.ReadAll(os.Stdin)
		return string(data), err
	}

	// TODO: adapter for clipboard package
	clipboard_data := clipboard.Read(clipboard.FmtText)
	if json.Valid(clipboard_data) {
		return string(clipboard_data), nil
	}

	return "", fmt.Errorf("No input provided")
}
