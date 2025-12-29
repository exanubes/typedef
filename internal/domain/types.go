package domain

import (
	"fmt"
	"slices"
	"strings"
)

type Type interface {
	Name() string
	Canonical() string
}

type ObjectType struct {
	Fields map[string]Type
}

func (t ObjectType) Name() string { return "object" }
func (t ObjectType) Canonical() string {
	fields := make([]string, 0)

	for key, field := range t.Fields {
		fields = append(fields, fmt.Sprintf("%s:%s", key, field.Canonical()))
	}
	slices.SortStableFunc(fields, func(a, b string) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}

		return 0
	})

	return fmt.Sprintf("object{%s}", strings.Join(fields, ","))
}

type ArrayType struct {
	Element Type
}

func (t ArrayType) Name() string { return "array" }
func (t ArrayType) Canonical() string {
	el := t.Element.Canonical()
	return fmt.Sprintf("array<%s>", el)
}

type BooleanType struct {
}

func (t BooleanType) Name() string      { return "boolean" }
func (t BooleanType) Canonical() string { return "boolean" }

type StringType struct {
}

func (t StringType) Name() string      { return "string" }
func (t StringType) Canonical() string { return "string" }

type IntType struct {
}

func (t IntType) Name() string      { return "int" }
func (t IntType) Canonical() string { return "int" }

type FloatType struct {
}

func (t FloatType) Name() string      { return "float" }
func (t FloatType) Canonical() string { return "float" }

type NullType struct{}

func (t NullType) Name() string      { return "null" }
func (t NullType) Canonical() string { return "null" }

type DateType struct{}

func (t DateType) Name() string      { return "date" }
func (t DateType) Canonical() string { return "date" }

type NamedType struct {
	Namespace string
	Identity  *ObjectType
	Hash      string
}

func (t NamedType) Name() string { return "named" }
func (t NamedType) Canonical() string {
	return fmt.Sprintf("named(%s#%s)", t.Namespace, t.Hash[:8])
}

type UnionType struct {
	OneOf []Type
}

func (t UnionType) Name() string { return "union" }
func (t UnionType) Canonical() string {
	names := make([]string, len(t.OneOf))
	for index, t := range t.OneOf {
		names[index] = t.Canonical()
	}
	slices.SortStableFunc(names, func(a string, b string) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})

	return fmt.Sprintf("union<%s>", strings.Join(names, "|"))
}

type UnknownType struct {
}

func (t UnknownType) Name() string      { return "unknown" }
func (t UnknownType) Canonical() string { return "unknown" }
