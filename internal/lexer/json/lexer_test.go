package json

import (
	"testing"

	"github.com/exanubes/typedef/internal/domain"
)

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
