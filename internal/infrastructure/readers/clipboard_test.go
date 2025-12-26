package readers

import (
	"testing"
)

func TestClipboardReader_Read(t *testing.T) {
	// Note: ClipboardReader depends on the golang.design/x/clipboard package which
	// directly accesses the system clipboard. Testing actual clipboard content requires
	// integration tests or mocking the clipboard package.
	//
	// The ClipboardReader validates that clipboard content is valid JSON using json.Valid().
	// If the clipboard is empty or contains non-JSON data, it returns an error.
	//
	// These tests verify the expected behavior when the system clipboard is in various states.
	// Actual clipboard functionality is better tested through integration tests at the CLI level.

	t.Run("clipboard reader can be instantiated", func(t *testing.T) {
		reader := NewClipboardReader()
		if reader == nil {
			t.Fatal("Expected non-nil ClipboardReader")
		}
	})

	t.Run("read returns error or valid result", func(t *testing.T) {
		// This test verifies that the Read method either:
		// 1. Returns valid JSON data from clipboard (success case)
		// 2. Returns "Clipboard is empty" error (when clipboard has no valid JSON)
		//
		// The actual result depends on the current system clipboard state.
		// In CI/CD environments, the clipboard is typically empty or unavailable.

		reader := NewClipboardReader()
		result, err := reader.Read()

		if err != nil {
			// Error case: clipboard is empty or contains non-JSON data
			expectedErrMsg := "Clipboard is empty"
			if err.Error() != expectedErrMsg {
				t.Fatalf("Expected error message %q, got %q", expectedErrMsg, err.Error())
			}
			if result != "" {
				t.Fatalf("Expected empty result on error, got %q", result)
			}
		} else {
			// Success case: clipboard contains valid JSON
			// We can't predict the exact content, but it should be non-empty
			if result == "" {
				t.Fatal("Expected non-empty result on success")
			}
		}
	})
}

// Integration test notes:
//
// To properly test ClipboardReader with specific clipboard content, you would need to:
// 1. Mock the clipboard.Read() function (requires refactoring to use an interface)
// 2. Use integration tests that set actual clipboard content before testing
// 3. Use a test helper that manipulates the system clipboard
//
// Example integration test approach:
//
// func TestClipboardReader_WithMockedClipboard(t *testing.T) {
//     // This would require refactoring ClipboardReader to accept a clipboard interface:
//     // type ClipboardService interface {
//     //     Read(format clipboard.Format) []byte
//     // }
//     //
//     // Then you could inject a mock:
//     // mockClipboard := &MockClipboard{
//     //     data: []byte(`{"test": "data"}`),
//     // }
//     // reader := NewClipboardReaderWithService(mockClipboard)
//     // result, err := reader.Read()
//     // // assert result == `{"test": "data"}` and err == nil
// }
