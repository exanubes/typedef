//go:build integration

package readers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestStdinReader_Integration tests StdinReader with actual piped stdin using subprocess pattern.
//
// The subprocess pattern allows us to test stdin piping by:
// 1. Parent test spawns a subprocess running the same test binary
// 2. Parent pipes test data to subprocess stdin via cmd.Stdin
// 3. Subprocess reads from stdin using StdinReader
// 4. Parent validates subprocess output matches expected input
//
// This pattern works in all environments including CI/CD.
//
// Run these tests with: go test -tags integration -v ./internal/infrastructure/readers/
func TestStdinReader_Integration(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple JSON object",
			input:    `{"name": "John", "age": 30}`,
			expected: `{"name": "John", "age": 30}`,
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
		},
		{
			name:     "large JSON payload",
			input:    generateLargeJSON(1024), // 1KB+ JSON
			expected: generateLargeJSON(1024),
		},
		{
			name:     "non-JSON text",
			input:    "This is plain text, not JSON",
			expected: "This is plain text, not JSON",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Subprocess mode: read from stdin and output result
			if os.Getenv("TEST_STDIN_SUBPROCESS") == "1" {
				reader := NewStdinReader()
				result, err := reader.Read()
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR: %v", err)
					os.Exit(1)
				}
				fmt.Print(result)
				os.Exit(0) // Exit immediately to avoid test framework output
			}

			// Parent mode: spawn subprocess with piped stdin
			cmd := exec.Command(os.Args[0], "-test.run=^"+t.Name()+"$")
			cmd.Env = append(os.Environ(), "TEST_STDIN_SUBPROCESS=1")
			cmd.Stdin = strings.NewReader(tc.input)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Subprocess failed: %v\nOutput: %s", err, output)
			}

			result := string(output)
			if result != tc.expected {
				t.Fatalf("Expected output %q, got %q", tc.expected, result)
			}
		})
	}
}

// generateLargeJSON creates a JSON string of approximately the specified size in bytes
func generateLargeJSON(sizeBytes int) string {
	// Create a JSON object with repeated fields to reach desired size
	var sb strings.Builder
	sb.WriteString(`{"data":[`)

	itemSize := len(`{"id":123,"value":"item"},`)
	itemCount := sizeBytes / itemSize

	for i := 0; i < itemCount; i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		fmt.Fprintf(&sb, `{"id":%d,"value":"item%d"}`, i, i)
	}

	sb.WriteString("]}")
	return sb.String()
}
