package transformer

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
