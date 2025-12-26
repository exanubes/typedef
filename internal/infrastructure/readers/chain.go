package readers

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type ChainReader struct {
	readers []domain.InputReader
}

func NewChainReader(readers ...domain.InputReader) *ChainReader {
	return &ChainReader{
		readers: readers,
	}
}

func (reader *ChainReader) Read() (string, error) {
	for _, reader := range reader.readers {
		if val, err := reader.Read(); err == nil {
			return val, nil
		}
	}

	return "", fmt.Errorf("No input provided")
}
