package readers

import (
	"encoding/json"
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type ClipboardReader struct {
	clipboard domain.Clipboard
}

func NewClipboardReader(clipboard domain.Clipboard) *ClipboardReader {
	return &ClipboardReader{
		clipboard: clipboard,
	}
}

func (reader *ClipboardReader) Read() (string, error) {
	clipboard_data, err := reader.clipboard.Read()

	if err != nil {
		return "", err
	}

	if json.Valid([]byte(clipboard_data)) {
		return string(clipboard_data), nil
	}

	return "", fmt.Errorf("Clipboard is empty")
}
