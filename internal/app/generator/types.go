package generator

import "github.com/exanubes/typedef/internal/app/ast"

type CodeGenerator interface {
	Generate(tree *ast.Program) string
}
