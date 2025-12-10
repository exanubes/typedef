package usecase_test

import (
	"errors"
	"testing"

	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/usecase"
)

type mockInputResolver struct {
	resolveFn func(domain.GenerateCommand) (domain.ResolvedInput, error)
	called    bool
}

func (m *mockInputResolver) Resolve(cmd domain.GenerateCommand) (domain.ResolvedInput, error) {
	m.called = true
	return m.resolveFn(cmd)
}

type mockCodegen struct {
	execFn func(domain.CodegenOptions) (string, error)
	called bool
}

func (m *mockCodegen) Execute(opts domain.CodegenOptions) (string, error) {
	m.called = true
	return m.execFn(opts)
}

type mockOutput struct {
	sendFn func(string, domain.OutputOptions) error
	called bool
}

func (m *mockOutput) Send(code string, out domain.OutputOptions) error {
	m.called = true
	return m.sendFn(code, out)
}

func TestGenerateUsecase_Run_Success(t *testing.T) {
	input := domain.ResolvedInput{
		Data:   string([]byte(`{"foo":"bar"}`)),
		Format: domain.TYPESCRIPT,
		Output: domain.OutputOptions{Target: "cli"},
	}

	inputMock := &mockInputResolver{
		resolveFn: func(cmd domain.GenerateCommand) (domain.ResolvedInput, error) {
			return input, nil
		},
	}

	codegenMock := &mockCodegen{
		execFn: func(opts domain.CodegenOptions) (string, error) {
			if opts.InputType != "json" {
				t.Fatalf("expected json input type")
			}
			return "generated code", nil
		},
	}

	outputMock := &mockOutput{
		sendFn: func(code string, out domain.OutputOptions) error {
			if code != "generated code" {
				t.Fatalf("unexpected code: %s", code)
			}
			return nil
		},
	}

	uc := usecase.NewGenerateUseCase(outputMock, inputMock, codegenMock)

	err := uc.Run(domain.GenerateCommand{})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !inputMock.called || !codegenMock.called || !outputMock.called {
		t.Fatalf("expected all ports to be called")
	}
}

func TestGenerateUsecase_Run_InputFailure(t *testing.T) {
	inputMock := &mockInputResolver{
		resolveFn: func(cmd domain.GenerateCommand) (domain.ResolvedInput, error) {
			return domain.ResolvedInput{}, errors.New("input resolver fail")
		},
	}

	uc := usecase.NewGenerateUseCase(nil, inputMock, nil)

	err := uc.Run(domain.GenerateCommand{})
	if err == nil {
		t.Fatal("expected error")
	}

	if !inputMock.called {
		t.Fatal("input resolver not called")
	}
}

func TestGenerateUsecase_Run_CodegenFailure(t *testing.T) {
	inputMock := &mockInputResolver{
		resolveFn: func(cmd domain.GenerateCommand) (domain.ResolvedInput, error) {
			return domain.ResolvedInput{
				Data:   string([]byte("{}")),
				Output: domain.OutputOptions{Target: "file"},
			}, nil
		},
	}

	codegenMock := &mockCodegen{
		execFn: func(opts domain.CodegenOptions) (string, error) {
			return "", errors.New("codegen fail")
		},
	}

	uc := usecase.NewGenerateUseCase(nil, inputMock, codegenMock)

	err := uc.Run(domain.GenerateCommand{})
	if err == nil {
		t.Fatal("expected codegen error")
	}
}

func TestGenerateUsecase_Run_OutputFailure(t *testing.T) {
	inputMock := &mockInputResolver{
		resolveFn: func(cmd domain.GenerateCommand) (domain.ResolvedInput, error) {
			return domain.ResolvedInput{
				Data:   string([]byte("{}")),
				Output: domain.OutputOptions{Target: "file"},
			}, nil
		},
	}

	codegenMock := &mockCodegen{
		execFn: func(opts domain.CodegenOptions) (string, error) {
			return "ok", nil
		},
	}

	outputMock := &mockOutput{
		sendFn: func(code string, out domain.OutputOptions) error {
			return errors.New("output fail")
		},
	}

	uc := usecase.NewGenerateUseCase(outputMock, inputMock, codegenMock)

	err := uc.Run(domain.GenerateCommand{})
	if err == nil {
		t.Fatal("expected output error")
	}
}
