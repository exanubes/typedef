package readers

import (
	"errors"
	"testing"

	"github.com/exanubes/typedef/internal/domain"
)

type MockInputReader struct {
	ReadFunc   func() (string, error)
	ReadCalled bool
}

func (m *MockInputReader) Read() (string, error) {
	m.ReadCalled = true
	if m.ReadFunc != nil {
		return m.ReadFunc()
	}
	return "", errors.New("mock not configured")
}

func TestChainReader_Read(t *testing.T) {
	testCases := []struct {
		name     string
		readers  []domain.InputReader
		expected string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "returns first successful reader result",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "first result", nil
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "second result", nil
				}},
			},
			expected: "first result",
			wantErr:  false,
		},
		{
			name: "tries readers in order until success",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("first failed")
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "second success", nil
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "third result", nil
				}},
			},
			expected: "second success",
			wantErr:  false,
		},
		{
			name: "returns error when all readers fail",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("first failed")
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("second failed")
				}},
			},
			wantErr: true,
			errMsg:  "No input provided",
		},
		{
			name:     "returns error for empty reader chain",
			readers:  []domain.InputReader{},
			wantErr:  true,
			errMsg:   "No input provided",
		},
		{
			name: "handles single reader success",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "single success", nil
				}},
			},
			expected: "single success",
			wantErr:  false,
		},
		{
			name: "handles single reader failure",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("single failed")
				}},
			},
			wantErr: true,
			errMsg:  "No input provided",
		},
		{
			name: "last reader succeeds after all others fail",
			readers: []domain.InputReader{
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("failed 1")
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "", errors.New("failed 2")
				}},
				&MockInputReader{ReadFunc: func() (string, error) {
					return "last success", nil
				}},
			},
			expected: "last success",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain := NewChainReader(tc.readers...)
			result, err := chain.Read()

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

func TestChainReader_ShortCircuits(t *testing.T) {
	firstReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "first", nil
	}}
	secondReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "second", nil
	}}
	thirdReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "third", nil
	}}

	chain := NewChainReader(firstReader, secondReader, thirdReader)
	result, err := chain.Read()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "first" {
		t.Fatalf("Expected result %q, got %q", "first", result)
	}

	if !firstReader.ReadCalled {
		t.Fatal("First reader should have been called")
	}

	if secondReader.ReadCalled {
		t.Fatal("Second reader should NOT have been called (short-circuit)")
	}

	if thirdReader.ReadCalled {
		t.Fatal("Third reader should NOT have been called (short-circuit)")
	}
}

func TestChainReader_CallsAllReadersUntilSuccess(t *testing.T) {
	firstReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "", errors.New("first failed")
	}}
	secondReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "", errors.New("second failed")
	}}
	thirdReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "third success", nil
	}}
	fourthReader := &MockInputReader{ReadFunc: func() (string, error) {
		return "fourth", nil
	}}

	chain := NewChainReader(firstReader, secondReader, thirdReader, fourthReader)
	result, err := chain.Read()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "third success" {
		t.Fatalf("Expected result %q, got %q", "third success", result)
	}

	if !firstReader.ReadCalled {
		t.Fatal("First reader should have been called")
	}

	if !secondReader.ReadCalled {
		t.Fatal("Second reader should have been called")
	}

	if !thirdReader.ReadCalled {
		t.Fatal("Third reader should have been called")
	}

	if fourthReader.ReadCalled {
		t.Fatal("Fourth reader should NOT have been called (short-circuit after third)")
	}
}
