package json

import (
	"fmt"
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

func (parser *AstParser) Parse() (*ast.Program, error) {
	program := &ast.Program{}
	for !parser.current_is(domain.EOF) {
		var value ast.Node
		var err error
		switch parser.current.Type {
		case domain.LBRACE:
			value, err = parser.parse_object()
		case domain.LBRACKET:
			value, err = parser.parse_array()
		default:
			return nil, fmt.Errorf("Expected { or [, received %s", parser.current.Literal)
		}

		if err != nil {
			return nil, err
		}

		program.Value = value
	}

	return program, nil
}

func (parser *AstParser) next_token() {
	parser.current = parser.next
	parser.next = parser.lexer.NextToken()
}

// Entry token: {
// Exit  token: one after }
func (parser *AstParser) parse_object() (*ast.ObjectNode, error) {
	if err := parser.expect(domain.LBRACE); err != nil {
		return nil, err
	}

	var result ast.ObjectNode

	if parser.advance_if(domain.RBRACE) {
		return &result, nil
	}

	if parser.current_is(domain.STRING) {
		node, err := parser.create_pair_node()
		if err != nil {
			return nil, err
		}

		result.Pairs = append(result.Pairs, node)

		for parser.advance_if(domain.COMMA) {
			node, err := parser.create_pair_node()
			if err != nil {
				return nil, err
			}
			result.Pairs = append(result.Pairs, node)
		}

		if err := parser.expect(domain.RBRACE); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Expected value, received %s", parser.current.Literal)
	}

	return &result, nil
}

// Entry token: string
// Exit  token: comma OR }
func (parser *AstParser) create_pair_node() (*ast.PairNode, error) {
	var result ast.PairNode
	val, err := parser.parse_value()

	if err != nil {
		return nil, err
	}

	result.Key = val.(*ast.StringNode)
	if err := parser.expect(domain.COLON); err != nil {
		return nil, err
	}
	result.Value, err = parser.parse_value()

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Entry token: [
// Exit  token: one after ]
func (parser *AstParser) parse_array() (*ast.ArrayNode, error) {
	var result ast.ArrayNode

	if err := parser.expect(domain.LBRACKET); err != nil {
		return nil, err
	}

	if parser.advance_if(domain.RBRACKET) {
		return &result, nil
	}
	value, err := parser.parse_value()
	if err != nil {
		return nil, err
	}
	result.Elements = append(result.Elements, value)

	for parser.advance_if(domain.COMMA) {
		value, err := parser.parse_value()
		if err != nil {
			return nil, err
		}
		result.Elements = append(result.Elements, value)
	}

	return &result, parser.expect(domain.RBRACKET)
}

// Enter token: value
// Exit  token: next token
func (parser *AstParser) parse_value() (ast.Node, error) {
	var result ast.Node
	var err error
	switch parser.current.Type {
	case domain.LBRACE:
		result, err = parser.parse_object()
	case domain.LBRACKET:
		result, err = parser.parse_array()
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
		err = fmt.Errorf("Expected value, received %s", parser.current.Literal)
	}

	return result, err
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

func (parser *AstParser) expect(token domain.TokenType) error {
	if !parser.current_is(token) {
		return fmt.Errorf("Expected %s, received %s", token, parser.current.Literal)
	}

	parser.next_token()

	return nil
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
