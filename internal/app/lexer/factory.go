package lexer

import (
	"strings"

	"github.com/exanubes/typedef/internal/app/lexer/json"
)

type LexerFactory struct{}

func (factory LexerFactory) Create(t, input string) Lexer {
	switch strings.ToLower(t) {
	case "json":
		return json.New(input)
	}
	return nil
}
