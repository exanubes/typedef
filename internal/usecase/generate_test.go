package usecase

import (
	"errors"
	"strings"
	"testing"

	"github.com/exanubes/typedef/internal/domain"
)

// MockCodegenService implements domain.CodegenService for testing
type MockCodegenService struct {
	ExecuteFunc   func(opts domain.CodegenOptions) (string, error)
	ExecuteCalled bool
	LastOptions   domain.CodegenOptions
}

func (m *MockCodegenService) Execute(opts domain.CodegenOptions) (string, error) {
	m.ExecuteCalled = true
	m.LastOptions = opts
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(opts)
	}
	return "", nil
}

func TestGenerateUsecase_Run_Success(t *testing.T) {
	expectedCode := "type User struct {\n\tName string\n}"
	mock := &MockCodegenService{
		ExecuteFunc: func(opts domain.CodegenOptions) (string, error) {
			return expectedCode, nil
		},
	}

	usecase := NewGenerateUseCase(mock)
	input := domain.GenerateCommandInput{
		Input:     `{"name": "John"}`,
		InputType: "json",
		Format:    domain.GOLANG,
	}

	output, err := usecase.Run(input)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if output.Code != expectedCode {
		t.Fatalf("Expected code %q, got: %q", expectedCode, output.Code)
	}

	if output.Format != domain.GOLANG {
		t.Fatalf("Expected format %v, got: %v", domain.GOLANG, output.Format)
	}

	if !mock.ExecuteCalled {
		t.Fatal("Expected Execute to be called")
	}
}

func TestGenerateUsecase_Run_ErrorPropagation(t *testing.T) {
	expectedError := errors.New("lexer failed to parse")
	mock := &MockCodegenService{
		ExecuteFunc: func(opts domain.CodegenOptions) (string, error) {
			return "", expectedError
		},
	}

	usecase := NewGenerateUseCase(mock)
	input := domain.GenerateCommandInput{
		Input:     `{"name": "John"}`,
		InputType: "json",
		Format:    domain.TYPESCRIPT,
	}

	output, err := usecase.Run(input)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "Failed execution of codegen service") {
		t.Fatalf("Expected error to contain context message, got: %v", err)
	}

	if !errors.Is(err, expectedError) {
		t.Fatalf("Expected error to wrap original error %v, got: %v", expectedError, err)
	}

	if output.Code != "" {
		t.Fatalf("Expected empty code on error, got: %q", output.Code)
	}
}

func TestGenerateUsecase_Run_InputMapping(t *testing.T) {
	mock := &MockCodegenService{
		ExecuteFunc: func(opts domain.CodegenOptions) (string, error) {
			return "generated code", nil
		},
	}

	usecase := NewGenerateUseCase(mock)
	input := domain.GenerateCommandInput{
		Input:     `{"user": {"id": 1, "name": "Alice"}}`,
		InputType: "json",
		Format:    domain.ZOD,
	}

	_, err := usecase.Run(input)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if mock.LastOptions.Input != input.Input {
		t.Fatalf("Expected Input %q, got: %q", input.Input, mock.LastOptions.Input)
	}

	if mock.LastOptions.InputType != input.InputType {
		t.Fatalf("Expected InputType %q, got: %q", input.InputType, mock.LastOptions.InputType)
	}

	if mock.LastOptions.Format != input.Format {
		t.Fatalf("Expected Format %v, got: %v", input.Format, mock.LastOptions.Format)
	}
}

func TestGenerateUsecase_Run_FormatPreservation(t *testing.T) {
	testCases := []struct {
		name   string
		format domain.Format
	}{
		{"golang format", domain.GOLANG},
		{"typescript format", domain.TYPESCRIPT},
		{"zod format", domain.ZOD},
		{"jsdoc format", domain.JSDOC},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &MockCodegenService{
				ExecuteFunc: func(opts domain.CodegenOptions) (string, error) {
					return "generated code", nil
				},
			}

			usecase := NewGenerateUseCase(mock)
			input := domain.GenerateCommandInput{
				Input:     `{"test": true}`,
				InputType: "json",
				Format:    tc.format,
			}

			output, err := usecase.Run(input)

			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			if output.Format != tc.format {
				t.Fatalf("Expected format %v, got: %v", tc.format, output.Format)
			}
		})
	}
}

func TestNewGenerateUseCase(t *testing.T) {
	mock := &MockCodegenService{}
	usecase := NewGenerateUseCase(mock)

	if usecase == nil {
		t.Fatal("Expected usecase to be non-nil")
	}

	if usecase.codegen != mock {
		t.Fatal("Expected usecase.codegen to be set to provided service")
	}
}
