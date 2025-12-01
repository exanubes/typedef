package graph

import (
	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/domain"
)

type TypeGraph interface {
	Generate(tree *ast.Program) *domain.ObjectType
}
