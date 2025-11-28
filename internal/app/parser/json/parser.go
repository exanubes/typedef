package parser

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
	for parser.current.Type != domain.EOF {
		var value ast.Node
		switch parser.current.Literal {
		case domain.LBRACE:
			parser.next_token()
			value = parser.parse_object()
		case domain.LBRACKET:
			parser.next_token()
			value = parser.parse_array()
		default:
			log.Fatalf("#Expected { or [, received %s", parser.current.Literal)
		}

		program.Value = value
	}

	return program
}

func (parser *AstParser) next_token() {
	parser.current = parser.next
	parser.next = parser.lexer.NextToken()
}

func (parser *AstParser) parse_object() ast.ObjectNode {
	var result ast.ObjectNode
	switch parser.current.Type {
	case domain.RBRACE:
		return result
	case domain.STRING:
		result.Pairs = append(result.Pairs, parser.create_pair_node())
		for parser.current.Type == domain.COMMA {
			parser.next_token()
			result.Pairs = append(result.Pairs, parser.create_pair_node())
		}
		if parser.current.Type != domain.RBRACE {
			log.Fatalf("Expected }, received %s", parser.current.Literal)
		}
		parser.next_token()
	default:
		log.Fatalf("Expected }, received %s", parser.current.Literal)
	}

	return result
}

func (parser *AstParser) create_pair_node() *ast.PairNode {
	var result ast.PairNode
	result.Key = parser.create_string_node()
	parser.next_token()
	if parser.current.Type != domain.COLON {
		log.Fatalf("Expected :, received %s", parser.current.Literal)
	}

	parser.next_token()
	switch parser.current.Type {
	case domain.LBRACE:
		parser.next_token()
		result.Value = parser.parse_object()
	case domain.LBRACKET:
		parser.next_token()
		result.Value = parser.parse_array()
	case domain.NUMBER:
		result.Value = parser.create_number_node()
	case domain.STRING:
		result.Value = parser.create_string_node()
	case domain.TRUE, domain.FALSE:
		result.Value = &ast.BooleanNode{Value: parser.current.Literal}
	case domain.NULL:
		result.Value = &ast.NullNode{}
	default:
		log.Fatalf("Unexpected token: %s", parser.current.Literal)
	}
	parser.next_token()
	return &result
}

func (parser *AstParser) parse_array() ast.ArrayNode {
	var result ast.ArrayNode
	return result
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
