//go:build integration

package readers

import (
	"testing"
	"time"

	"github.com/exanubes/typedef/internal/infrastructure/clipboard"
)

// This test:
// 1. Creates a clipboard instance using clipboard.New()
// 2. If clipboard is unavailable (headless/CI/CD environment), skip test gracefully
// 3. Write test data to clipboard
// 4. Read with ClipboardReader and validate result
//
// The 10ms delay after clipboard.Write() ensures the clipboard is updated before reading.
//
// Run these tests with: go test -tags integration -v ./internal/infrastructure/readers/
//
// Note: These tests may skip in CI/CD environments where clipboard is unavailable.
// This is expected behavior - tests skip instead of failing.
func TestClipboardReader_Integration(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid JSON object",
			input:    `{"name": "John", "age": 30}`,
			expected: `{"name": "John", "age": 30}`,
			wantErr:  false,
		},
		{
			name: "multi-line JSON",
			input: `{
  "name": "Jane",
  "age": 25,
  "city": "New York"
}`,
			expected: `{
  "name": "Jane",
  "age": 25,
  "city": "New York"
}`,
			wantErr: false,
		},
		{
			name:    "non-JSON text",
			input:   "This is plain text, not JSON",
			wantErr: true,
			errMsg:  "Clipboard is empty",
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
			errMsg:  "Clipboard is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cb := clipboard.New()
			if cb == nil {
				t.Skip("Clipboard unavailable (this is expected in CI/CD environments)")
				return
			}

			if err := cb.Write(tc.input); err != nil {
				t.Skipf("Failed to write to clipboard: %v", err)
				return
			}

			// NOTE: Small delay to ensure clipboard is updated
			time.Sleep(10 * time.Millisecond)

			reader := NewClipboardReader(cb)
			result, err := reader.Read()

			if tc.wantErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				if err.Error() != tc.errMsg {
					t.Fatalf("Expected error message %q, got %q", tc.errMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tc.expected {
				t.Fatalf("Expected result %q, got %q", tc.expected, result)
			}
		})
	}
}
