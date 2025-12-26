package readers

import "fmt"

type FlagReader struct {
	value string
}

func NewFlagReader(value string) *FlagReader {
	return &FlagReader{
		value: value,
	}
}

func (reader *FlagReader) Read() (string, error) {
	if reader.value == "" {
		return "", fmt.Errorf("flag value is empty")
	}

	return reader.value, nil
}
