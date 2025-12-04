package transformer

import (
	"github.com/exanubes/typedef/internal/domain"
)

type TypeDef struct {
	ID     string
	Kind   TypeKind
	Name   string
	Fields []FieldDef
}

type FieldDef struct {
	Name       string
	ParsedName string
	TypeID     string
}

type TypeKind string

const (
	KindObject = "OBJECT"
)

type Transformer interface {
	Transform(tree *domain.ObjectType) []TypeDef
}
