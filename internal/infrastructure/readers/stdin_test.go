//go:build !integration

package readers

import (
	"testing"
)

func TestStdinReader_Read(t *testing.T) {
	// Note: This test only covers the error case when no stdin is piped.
	// Testing actual piped stdin input requires subprocess testing or integration tests,
	// which is better handled at the CLI driver level.
	//
	// The StdinReader checks if stdin has piped data by examining the file mode
	// using os.Stdin.Stat(). In a normal test environment without piped input,
	// this will return an error indicating stdin is empty.

	t.Run("returns error when no stdin is piped", func(t *testing.T) {
		reader := NewStdinReader()
		_, err := reader.Read()

		if err == nil {
			t.Fatal("Expected error when no stdin is piped, got nil")
		}

		expectedErrMsg := "Stdin is empty"
		if err.Error() != expectedErrMsg {
			t.Fatalf("Expected error message %q, got %q", expectedErrMsg, err.Error())
		}
	})
}

// Integration test example (commented out):
// To test actual piped stdin, you would need to create a subprocess test like:
//
// func TestStdinReader_ReadPipedInput(t *testing.T) {
//     if os.Getenv("TEST_STDIN_SUBPROCESS") == "1" {
//         reader := NewStdinReader()
//         result, err := reader.Read()
//         if err != nil {
//             fmt.Fprintf(os.Stderr, "ERROR: %v", err)
//             os.Exit(1)
//         }
//         fmt.Print(result)
//         return
//     }
//
//     cmd := exec.Command(os.Args[0], "-test.run=TestStdinReader_ReadPipedInput")
//     cmd.Env = append(os.Environ(), "TEST_STDIN_SUBPROCESS=1")
//     cmd.Stdin = strings.NewReader("test input")
//     output, err := cmd.CombinedOutput()
//     if err != nil {
//         t.Fatalf("Subprocess failed: %v\nOutput: %s", err, output)
//     }
//     if string(output) != "test input" {
//         t.Fatalf("Expected %q, got %q", "test input", string(output))
//     }
// }
