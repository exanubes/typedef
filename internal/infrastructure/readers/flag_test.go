//go:build !integration

package readers

import (
	"testing"
)

func TestFlagReader_Read(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid non-empty value",
			value:    `{"name": "John", "age": 30}`,
			expected: `{"name": "John", "age": 30}`,
			wantErr:  false,
		},
		{
			name:     "simple string value",
			value:    "test input",
			expected: "test input",
			wantErr:  false,
		},
		{
			name:     "multi-line value",
			value:    "line1\nline2\nline3",
			expected: "line1\nline2\nline3",
			wantErr:  false,
		},
		{
			name:     "whitespace-only value",
			value:    "   \t\n   ",
			expected: "   \t\n   ",
			wantErr:  false,
		},
		{
			name:    "empty string value",
			value:   "",
			wantErr: true,
			errMsg:  "flag value is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := NewFlagReader(tc.value)
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
