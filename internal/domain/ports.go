package domain

type CodegenService interface {
	Execute(CodegenOptions) (string, error)
}

type CodegenOptions struct {
	InputType  string
	OutputType string
	Input      string
	Format     int // TODO: is format part of domain?
}

type OutputTarget interface {
	Send(string) error
}
