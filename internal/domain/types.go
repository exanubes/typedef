package domain

type Type interface {
	Name() string
}

type ObjectType struct {
	Fields map[string]Type
}

func (t ObjectType) Name() string { return "object" }

type ArrayType struct {
	Element Type
}

func (t ArrayType) Name() string { return "array" }

type BooleanType struct {
}

func (t BooleanType) Name() string { return "boolean" }

type StringType struct {
}

func (t StringType) Name() string { return "string" }

type IntType struct {
}

func (t IntType) Name() string { return "int" }

type FloatType struct {
}

func (t FloatType) Name() string { return "float" }

type NullType struct{}

func (t NullType) Name() string { return "null" }

type DateType struct{}

func (t DateType) Name() string { return "date" }

type NamedType struct {
	Namespace string
	Identity  Type
	Hash      string
}

func (t NamedType) Compare(input NamedType) bool {
	return true
}

func (t NamedType) Name() string { return "named" }

type UnionType struct {
	OneOf []Type
}

func (t UnionType) Name() string { return "union" }

type UnknownType struct {
}

func (t UnknownType) Name() string { return "unknown" }
