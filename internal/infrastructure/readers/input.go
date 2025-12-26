package readers

type InputReader struct{}

func NewInputReader() *InputReader {
	return &InputReader{}
}

func Read() (string, error) {
	return "", nil
}
