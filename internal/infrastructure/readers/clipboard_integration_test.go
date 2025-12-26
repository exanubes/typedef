//go:build integration

package readers

import (
	"testing"
	"time"

	"golang.design/x/clipboard"
)

// TestClipboardReader_Integration tests ClipboardReader with real clipboard access.
//
// This uses a hybrid approach:
// 1. Initialize clipboard with clipboard.Init()
// 2. If init fails (headless/CI/CD environment), skip test gracefully
// 3. If init succeeds, write test data to clipboard
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
			// Initialize clipboard - skip if unavailable
			if err := initClipboard(t); err != nil {
				t.Skipf("Clipboard unavailable: %v (this is expected in CI/CD environments)", err)
				return
			}

			// Write test data to clipboard
			clipboard.Write(clipboard.FmtText, []byte(tc.input))

			// Small delay to ensure clipboard is updated
			time.Sleep(10 * time.Millisecond)

			// Read with ClipboardReader
			reader := NewClipboardReader()
			result, err := reader.Read()

			// Validate result
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

// initClipboard initializes the clipboard and returns an error if it fails.
// This helper allows tests to gracefully skip when clipboard is unavailable.
func initClipboard(t *testing.T) error {
	// clipboard.Init() may panic in some environments, so we recover
	defer func() {
		if r := recover(); r != nil {
			t.Skipf("Clipboard initialization panicked: %v", r)
		}
	}()

	err := clipboard.Init()
	return err
}
