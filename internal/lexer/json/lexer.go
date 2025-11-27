package json

import (
	"github.com/exanubes/typedef/internal/domain"
)

type Lexer struct {
	input         string
	position      int
	read_position int
	character     byte
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.read_next_character()
	return lexer
}

func (lexer *Lexer) NextToken() domain.Token {
	var token domain.Token

	lexer.skip_whitespace()
	switch lexer.character {
	case ':':
		token = domain.NewToken(domain.COLON, lexer.character)
	case '{':
		token = domain.NewToken(domain.LBRACE, lexer.character)
	case '}':
		token = domain.NewToken(domain.RBRACE, lexer.character)
	case '[':
		token = domain.NewToken(domain.LBRACKET, lexer.character)
	case ']':
		token = domain.NewToken(domain.RBRACKET, lexer.character)
	case ',':
		token = domain.NewToken(domain.COMMA, lexer.character)
	case '-':
		token = domain.NewToken(domain.MINUS, lexer.character)
	case '"':
		token.Type = domain.STRING
		token.Literal = lexer.read_string()
	case 0:
		token.Literal = ""
		token.Type = domain.EOF
	default:
		if lexer.is_letter(lexer.character) {
			token.Literal = lexer.read_identifier()
			token.Type = domain.LookupIdentifier(token.Literal)
			return token
		} else if lexer.is_digit(lexer.character) {
			token.Literal = lexer.read_number()
			token.Type = domain.NUMBER
			return token
		} else {
			token = domain.NewToken(domain.ILLEGAL, lexer.character)
		}
	}

	lexer.read_next_character()

	return token
}

func (lexer *Lexer) skip_whitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.read_next_character()
	}
}

func (lexer *Lexer) read_next_character() {
	if lexer.read_position >= len(lexer.input) {
		lexer.character = 0 // zero is ascii code for NUL character
	} else {
		lexer.character = lexer.input[lexer.read_position]
	}

	lexer.position = lexer.read_position
	lexer.read_position += 1
}

func (lexer *Lexer) read_identifier() string {
	position := lexer.position
	for lexer.is_letter(lexer.character) {
		lexer.read_next_character()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) is_letter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func (lexer *Lexer) read_string() string {
	position := lexer.position + 1
	for {
		lexer.read_next_character()
		if lexer.character == '\\' {
			lexer.read_next_character()
			lexer.read_next_character()
		}
		if lexer.character == '"' || lexer.character == 0 {
			break
		}
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) is_digit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) read_number() string {
	position := lexer.position
	isFloat := false
	isExponential := false
	for lexer.is_digit(lexer.character) || !isFloat && lexer.character == '.' || !isExponential && (lexer.character == 'e' || lexer.character == 'E') {
		lexer.read_next_character()
	}

	return lexer.input[position:lexer.position]
}
