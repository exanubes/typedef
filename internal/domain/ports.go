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

type OutputOptions struct {
	Target string
	Path   string
}

type Clipboard interface {
	Read() (string, error)
	Write(input string) error
}
