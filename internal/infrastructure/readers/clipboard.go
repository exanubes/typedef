package readers

import (
	"encoding/json"
	"fmt"

	"golang.design/x/clipboard"
)

type ClipboardReader struct{}

func NewClipboardReader() *ClipboardReader {
	return &ClipboardReader{}
}

func (reader *ClipboardReader) Read() (string, error) {
	clipboard_data := clipboard.Read(clipboard.FmtText)

	if json.Valid(clipboard_data) {
		return string(clipboard_data), nil
	}

	return "", fmt.Errorf("Clipboard is empty")
}
