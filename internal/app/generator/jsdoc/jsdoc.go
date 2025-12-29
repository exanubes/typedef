package jsdoc

import (
	"fmt"
	"slices"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/utils"
)

type JSDocCodegen struct{}

func New() *JSDocCodegen {
	return &JSDocCodegen{}
}

func (generator *JSDocCodegen) Generate(root domain.Type) string {
	code := &strings.Builder{}
	visited := map[string]string{}
	return generator.dfs(root, visited, code)
}

func (generator *JSDocCodegen) dfs(node domain.Type, visited map[string]string, code *strings.Builder) string {
	id := node.Canonical()
	if val, ok := visited[id]; ok {
		return val
	}

	switch node := node.(type) {
	case *domain.ObjectType:
		builder := new_typedef_builder()
		builder.with_name("Root")
		for key, field := range node.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}
		code.WriteString(builder.build())

		return code.String()

	case *domain.NamedType:
		type_name := utils.Capitalize(node.Namespace)

		builder := new_typedef_builder()
		builder.with_name(type_name)

		for key, field := range node.Identity.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}

		code.WriteString(builder.build())
		code.WriteRune('\n')
		visited[id] = type_name

		return type_name

	case *domain.UnionType:
		types := []string{}
		for _, member := range node.OneOf {
			types = append(types, generator.dfs(member, visited, code))
		}
		slices.SortStableFunc(types, func(a, b string) int {
			if a > b {
				return 1
			}

			if a < b {
				return -1
			}

			return 0
		})

		union_type := strings.Join(types, " | ")
		visited[id] = union_type
		return union_type
	case *domain.ArrayType:
		field_type := generator.dfs(node.Element, visited, code)
		_, is_union := node.Element.(*domain.UnionType)
		var type_def string
		if is_union {
			type_def = fmt.Sprintf("(%s)[]", field_type)
		} else {
			type_def = fmt.Sprintf("%s[]", field_type)
		}
		visited[id] = type_def
		return type_def
	case *domain.IntType, *domain.FloatType, domain.IntType, domain.FloatType:
		visited[id] = "number"
		return "number"
	case *domain.NullType, domain.NullType:
		visited[id] = "null"
		return "null"
	case *domain.UnknownType, domain.UnknownType:
		visited[id] = "unknown"
		return "unknown"
	case *domain.DateType, domain.DateType:
		visited[id] = "Date"
		return "Date"
	case *domain.StringType, domain.StringType:
		visited[id] = "string"
		return "string"
	case *domain.BooleanType, domain.BooleanType:
		visited[id] = "boolean"
		return "boolean"
	case *domain.UuidType, domain.UuidType:
		visited[id] = "string"
		// TODO: make configurable to use external type e.g., @types/uuid
		return "string"
	default:
		return fmt.Sprintf("Unhandled type '%s'", node.Canonical())
	}

}
