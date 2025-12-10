package domain

type CodegenService interface {
	Execute(CodegenOptions) (string, error)
}

type CodegenOptions struct {
	InputType  string
	OutputType string
	Input      string
	Format     Format
}

type OutputTarget interface {
	Send(string) error
}

type OutputService interface {
	Send(payload string, options OutputOptions) error
}

type InputResolver interface {
	Resolve(input GenerateCommand) (ResolvedInput, error)
}

type ResolvedInput struct {
	Data   string
	Output OutputOptions
	Format Format
}

type OutputOptions struct {
	Target string
	Path   string
}
