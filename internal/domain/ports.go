package domain

type CodegenService interface {
	Execute(CodegenOptions) (string, error)
}

type CodegenOptions struct {
	InputType  string
	OutputType string
	Input      string
}

type OutputTarget interface {
	Send(string) error
}
