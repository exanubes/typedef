package json

import (
	"log"
	"strings"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/lexer"
	"github.com/exanubes/typedef/internal/domain"
)

type AstParser struct {
	lexer   lexer.Lexer
	current domain.Token
	next    domain.Token
}

func New(lexer lexer.Lexer) *AstParser {
	parser := &AstParser{
		lexer: lexer,
	}

	// NOTE: setting current and next to be tokens
	parser.next_token()
	parser.next_token()

	return parser
}

func (parser *AstParser) Parse() *ast.Program {
	program := &ast.Program{}
	for !parser.current_is(domain.EOF) {
		var value ast.Node
		switch parser.current.Type {
		case domain.LBRACE:
			value = parser.parse_object()
		case domain.LBRACKET:
			value = parser.parse_array()
		default:
			log.Fatalf("Expected { or [, received %s", parser.current.Literal)
		}

		program.Value = value
	}

	return program
}

func (parser *AstParser) next_token() {
	parser.current = parser.next
	parser.next = parser.lexer.NextToken()
}

// Entry token: {
// Exit  token: one after }
func (parser *AstParser) parse_object() *ast.ObjectNode {
	parser.expect(domain.LBRACE)

	var result ast.ObjectNode

	if parser.advance_if(domain.RBRACE) {
		return &result
	}

	if parser.current_is(domain.STRING) {
		result.Pairs = append(result.Pairs, parser.create_pair_node())

		for parser.advance_if(domain.COMMA) {
			result.Pairs = append(result.Pairs, parser.create_pair_node())
		}

		parser.expect(domain.RBRACE)
	} else {
		log.Fatalf("Expected value, received %s", parser.current.Literal)
	}

	return &result
}

// Entry token: string
// Exit  token: comma OR }
func (parser *AstParser) create_pair_node() *ast.PairNode {
	var result ast.PairNode
	result.Key = parser.parse_value().(*ast.StringNode)
	parser.expect(domain.COLON)
	result.Value = parser.parse_value()
	return &result
}

// Entry token: [
// Exit  token: one after ]
func (parser *AstParser) parse_array() *ast.ArrayNode {
	var result ast.ArrayNode
	parser.expect(domain.LBRACKET)

	if parser.advance_if(domain.RBRACKET) {
		return &result
	}

	result.Elements = append(result.Elements, parser.parse_value())

	for parser.advance_if(domain.COMMA) {
		result.Elements = append(result.Elements, parser.parse_value())
	}

	parser.expect(domain.RBRACKET)

	return &result
}

// Enter token: value
// Exit  token: next token
func (parser *AstParser) parse_value() ast.Node {
	var result ast.Node
	switch parser.current.Type {
	case domain.LBRACE:
		result = parser.parse_object()
	case domain.LBRACKET:
		result = parser.parse_array()
	case domain.NUMBER:
		result = parser.create_number_node()
		parser.next_token()
	case domain.STRING:
		result = parser.create_string_node()
		parser.next_token()
	case domain.TRUE, domain.FALSE:
		result = parser.create_boolean_node()
		parser.next_token()
	case domain.NULL:
		result = parser.create_null_node()
		parser.next_token()
	default:
		log.Fatalf("Expected value, received %s", parser.current.Literal)
	}

	return result
}

func (parser *AstParser) create_boolean_node() *ast.BooleanNode {
	var result ast.BooleanNode
	result.Value = parser.current.Literal
	return &result
}

func (parser *AstParser) create_null_node() *ast.NullNode {
	return &ast.NullNode{}
}

func (parser *AstParser) create_string_node() *ast.StringNode {
	var result ast.StringNode
	result.Value = parser.current.Literal
	return &result
}

func (parser *AstParser) create_number_node() *ast.NumberNode {
	var result ast.NumberNode
	result.Value = parser.current.Literal

	if strings.Contains(result.Value, ".") || strings.Contains(strings.ToLower(result.Value), "e") {
		result.Kind = ast.FLOAT
	} else {
		result.Kind = ast.INTEGER
	}

	return &result
}

func (parser *AstParser) expect(token domain.TokenType) {
	if !parser.current_is(token) {
		log.Fatalf("Expected %s, received %s", token, parser.current.Literal)
	}

	parser.next_token()
}

func (parser *AstParser) current_is(token domain.TokenType) bool {
	return parser.current.Type == token
}

func (parser *AstParser) advance_if(token domain.TokenType) bool {
	if parser.current_is(token) {
		parser.next_token()
		return true
	}

	return false
}
