//go:build !integration

package json

import (
	"strings"
	"testing"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/lexer/json"
)

func TestNextToken(test *testing.T) {
	input := `
	[{
	"property": true,
	"prop1": null
	}]`

	lexer := json.New(input)
	parser := New(lexer)

	program, _ := parser.Parse()

	array, ok := program.Value.(*ast.ArrayNode)

	if !ok {
		test.Fatalf("expect array, received %+v", array)
	}

	if len(array.Elements) != 1 {
		test.Fatalf("expected array to have 1 element, received %d", len(array.Elements))
	}

	object, ok := array.Elements[0].(*ast.ObjectNode)

	if !ok {
		test.Fatalf("expected array[0] to be an object, received %+v", object)
	}

	if len(object.Pairs) != 2 {
		test.Fatalf("expected object to have 2 key/value pairs, received %d", len(object.Pairs))
	}

	if object.Pairs[0].Key.Literal() != "property" {
		test.Fatalf("expected property, received %s", object.Pairs[0].Key.Literal())
	}

	if object.Pairs[0].Value.Literal() != "true" {
		test.Fatalf("expected true, received %s", object.Pairs[0].Value.Literal())
	}

	if object.Pairs[1].Key.Literal() != "prop1" {
		test.Fatalf("expected property, received %s", object.Pairs[0].Key.Literal())
	}

	if object.Pairs[1].Value.Literal() != "null" {
		test.Fatalf("expected true, received %s", object.Pairs[0].Value.Literal())
	}

}

// Helper Functions

func createTestParser(input string) *AstParser {
	lexer := json.New(input)
	return New(lexer)
}

func assertErrorContains(t *testing.T, err error, substring string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error but got nil")
	}
	if !strings.Contains(err.Error(), substring) {
		t.Fatalf("expected error to contain %q, got %q", substring, err.Error())
	}
}

// Error Path Tests

func TestParseTopLevelErrors(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "invalid_top_level_number",
			input:         `123`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_string",
			input:         `"string"`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_true",
			input:         `true`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_false",
			input:         `false`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_null",
			input:         `null`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_colon",
			input:         `:`,
			expectedError: "Expected { or [",
		},
		{
			name:          "invalid_top_level_closing_bracket",
			input:         `]`,
			expectedError: "Expected { or [",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := createTestParser(tt.input)
			program, err := parser.Parse()

			assertErrorContains(t, err, tt.expectedError)

			if program != nil {
				t.Fatal("expected program to be nil on error")
			}
		})
	}
}

func TestParseObjectErrors(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "missing_closing_brace",
			input:         `{"key": 1`,
			expectedError: "Expected }",
		},
		{
			name:          "missing_closing_brace_after_key",
			input:         `{"key"`,
			expectedError: "Expected :",
		},
		{
			name:          "missing_colon",
			input:         `{"key" 1}`,
			expectedError: "Expected :",
		},
		{
			name:          "missing_value_after_colon",
			input:         `{"key":}`,
			expectedError: "Expected value",
		},
		{
			name:          "invalid_object_key_number",
			input:         `{123: "value"}`,
			expectedError: "Expected value",
		},
		{
			name:          "invalid_object_key_boolean",
			input:         `{true: "value"}`,
			expectedError: "Expected value",
		},
		{
			name:          "invalid_object_key_array",
			input:         `{[]: "value"}`,
			expectedError: "Expected value",
		},
		{
			name:          "invalid_object_key_object",
			input:         `{{}: "value"}`,
			expectedError: "Expected value",
		},
		{
			name:          "trailing_comma_in_object",
			input:         `{"key": 1,}`,
			expectedError: "Expected value",
		},
		{
			name:          "missing_comma_between_pairs",
			input:         `{"key1": 1 "key2": 2}`,
			expectedError: "Expected }",
		},
		{
			name:          "eof_after_opening_brace",
			input:         `{`,
			expectedError: "Expected value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := createTestParser(tt.input)
			program, err := parser.Parse()

			assertErrorContains(t, err, tt.expectedError)

			if program != nil {
				t.Fatal("expected program to be nil on error")
			}
		})
	}
}

func TestParseArrayErrors(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "missing_closing_bracket",
			input:         `[1, 2`,
			expectedError: "Expected ]",
		},
		{
			name:          "missing_comma_between_elements",
			input:         `[1 2]`,
			expectedError: "Expected ]",
		},
		{
			name:          "trailing_comma_in_array",
			input:         `[1,]`,
			expectedError: "Expected value",
		},
		{
			name:          "unclosed_object_in_array",
			input:         `[{`,
			expectedError: "Expected value",
		},
		{
			name:          "unclosed_nested_array",
			input:         `[[`,
			expectedError: "Expected value",
		},
		{
			name:          "eof_after_opening_bracket",
			input:         `[`,
			expectedError: "Expected value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := createTestParser(tt.input)
			program, err := parser.Parse()

			assertErrorContains(t, err, tt.expectedError)

			if program != nil {
				t.Fatal("expected program to be nil on error")
			}
		})
	}
}

func TestParseNestedStructureErrors(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "nested_object_unclosed",
			input:         `{"user": {"name": "John"}`,
			expectedError: "Expected }",
		},
		{
			name:          "mismatched_brackets_array_to_object",
			input:         `{"items": [{"id": 1]}}`,
			expectedError: "Expected }",
		},
		{
			name:          "deeply_nested_array_unclosed",
			input:         `[[[[`,
			expectedError: "Expected value",
		},
		{
			name:          "mixed_nesting_error",
			input:         `[{"x": 1}, {"y": 2]`,
			expectedError: "Expected }",
		},
		{
			name:          "nested_object_missing_colon",
			input:         `{"outer": {"inner" "value"}}`,
			expectedError: "Expected :",
		},
		{
			name:          "array_with_nested_unclosed_object",
			input:         `[1, {"key": "value"`,
			expectedError: "Expected }",
		},
		{
			name:          "object_with_nested_unclosed_array",
			input:         `{"data": [1, 2, 3}`,
			expectedError: "Expected ]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := createTestParser(tt.input)
			program, err := parser.Parse()

			assertErrorContains(t, err, tt.expectedError)

			if program != nil {
				t.Fatal("expected program to be nil on error")
			}
		})
	}
}
