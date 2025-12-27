package readers

import (
	"testing"

	"github.com/exanubes/typedef/internal/infrastructure/clipboard"
)

func TestClipboardReader_Read(t *testing.T) {

	t.Run("clipboard reader can be instantiated", func(t *testing.T) {
		reader := NewClipboardReader(clipboard.New())
		if reader == nil {
			t.Fatal("Expected non-nil ClipboardReader")
		}
	})

	t.Run("read returns error or valid result", func(t *testing.T) {
		reader := NewClipboardReader(clipboard.New())
		result, err := reader.Read()

		if err != nil {
			expectedErrMsg := "Clipboard is empty"
			if err.Error() != expectedErrMsg {
				t.Fatalf("Expected error message %q, got %q", expectedErrMsg, err.Error())
			}
			if result != "" {
				t.Fatalf("Expected empty result on error, got %q", result)
			}
		} else {
			if result == "" {
				t.Fatal("Expected non-empty result on success")
			}
		}
	})
}
