package parser

import (
	"strings"

	"github.com/exanubes/typedef/internal/app/lexer"
	"github.com/exanubes/typedef/internal/app/parser/json"
)

type ParserFactory struct{}

func (factory ParserFactory) Create(t string, lexer lexer.Lexer) Parser {
	switch strings.ToLower(t) {
	case "json":
		return json.New(lexer)
	}
	return nil
}
