package parser

import (
	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/lexer"
)

type Parser interface {
	Parse() *ast.Program
}

type Factory interface {
	Create(string, lexer.Lexer) Parser
}
