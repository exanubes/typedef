package json

import (
	"testing"

	"github.com/exanubes/typedef/internal/domain"
)

func TestNumberParsing(test *testing.T) {
	testCases := []struct {
		input    string
		expected string
		name     string
	}{
		{"123", "123", "simple integer"},
		{"0", "0", "zero"},
		{"3.14159", "3.14159", "decimal number"},
		{"1e10", "1e10", "exponential notation"},
		{"1E10", "1E10", "exponential notation uppercase"},
		{"1e+5", "1e+5", "exponential with positive sign"},
		{"1e-5", "1e-5", "exponential with negative sign"},
		{"123.456e-10", "123.456e-10", "complex exponential"},
		{"0.5", "0.5", "decimal less than one"},
		{"42E+3", "42E+3", "exponential with plus sign uppercase"},
	}

	for _, tc := range testCases {
		test.Run(tc.name, func(t *testing.T) {
			lexer := New(tc.input)
			token := lexer.NextToken()

			if token.Type != domain.NUMBER {
				t.Errorf("Expected token type NUMBER, got %v", token.Type)
			}

			if token.Literal != tc.expected {
				t.Errorf("Expected literal %q, got %q", tc.expected, token.Literal)
			}
		})
	}
}

func TestMalformedNumbers(test *testing.T) {
	testCases := []struct {
		input       string
		expectedLit string
		name        string
	}{
		{"1.", "1.", "decimal without fractional part"},
		{"1e", "1e", "exponential without exponent"},
		{"1e+", "1e+", "exponential with sign but no digits"},
		{"1E-", "1E-", "exponential with negative sign but no digits"},
	}

	for _, tc := range testCases {
		test.Run(tc.name, func(t *testing.T) {
			lexer := New(tc.input)
			token := lexer.NextToken()

			if token.Type != domain.ILLEGAL {
				t.Errorf("Expected token type ILLEGAL, got %v", token.Type)
			}

			if token.Literal != tc.expectedLit {
				t.Errorf("Expected literal %q, got %q", tc.expectedLit, token.Literal)
			}
		})
	}
}

func TestNextToken(test *testing.T) {
	input := `
	{
	"boolean": true,
	"array": ["item1", "item2"],
	"no_value": null,
	"integer": 69420,
	"negative": -69420,
	"string_spaced": "hello world",
	"string_escaped": "hello \"escaped\" world",
	"float": 3.14159,
	"exponential": 1e10,
	"neg_float": -12.75
	}`

	tests := []struct {
		expectedType    domain.TokenType
		expectedLiteral string
	}{
		{domain.LBRACE, "{"},
		{domain.STRING, "boolean"},
		{domain.COLON, ":"},
		{domain.TRUE, "true"},
		{domain.COMMA, ","},
		{domain.STRING, "array"},
		{domain.COLON, ":"},
		{domain.LBRACKET, "["},
		{domain.STRING, "item1"},
		{domain.COMMA, ","},
		{domain.STRING, "item2"},
		{domain.RBRACKET, "]"},
		{domain.COMMA, ","},
		{domain.STRING, "no_value"},
		{domain.COLON, ":"},
		{domain.NULL, "null"},
		{domain.COMMA, ","},
		{domain.STRING, "integer"},
		{domain.COLON, ":"},
		{domain.NUMBER, "69420"},
		{domain.COMMA, ","},
		{domain.STRING, "negative"},
		{domain.COLON, ":"},
		{domain.MINUS, "-"},
		{domain.NUMBER, "69420"},
		{domain.COMMA, ","},
		{domain.STRING, "string_spaced"},
		{domain.COLON, ":"},
		{domain.STRING, "hello world"},
		{domain.COMMA, ","},
		{domain.STRING, "string_escaped"},
		{domain.COLON, ":"},
		{domain.STRING, "hello \\\"escaped\\\" world"}, // literal should match lexer output
		{domain.COMMA, ","},
		{domain.STRING, "float"},
		{domain.COLON, ":"},
		{domain.NUMBER, "3.14159"},
		{domain.COMMA, ","},
		{domain.STRING, "exponential"},
		{domain.COLON, ":"},
		{domain.NUMBER, "1e10"},
		{domain.COMMA, ","},
		{domain.STRING, "neg_float"},
		{domain.COLON, ":"},
		{domain.MINUS, "-"},
		{domain.NUMBER, "12.75"},
		{domain.RBRACE, "}"},
		{domain.EOF, ""},
	}

	lexer := New(input)

	for index, t := range tests {
		token := lexer.NextToken()
		if token.Type != t.expectedType {
			test.Fatalf("tests[%d] - wrong TokenType, expected=%q, received=%q", index, t.expectedType, token.Type)
		}

		if token.Literal != t.expectedLiteral {
			test.Fatalf("tests[%d] - wrong Literal, expected=%q, received=%q", index, t.expectedLiteral, token.Literal)
		}
	}
}
