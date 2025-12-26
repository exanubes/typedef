package readers

import (
	"fmt"
	"io"
	"os"
)

type StdinReader struct{}

func NewStdinReader() *StdinReader {
	return &StdinReader{}
}

func (reader *StdinReader) Read() (string, error) {
	stat, _ := os.Stdin.Stat()
	has_piped_input := (stat.Mode() & os.ModeCharDevice) == 0

	if has_piped_input {
		data, err := io.ReadAll(os.Stdin)
		return string(data), err
	}

	return "", fmt.Errorf("Stdin is empty")
}
