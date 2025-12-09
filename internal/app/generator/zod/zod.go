package zod

import (
	"fmt"
	"slices"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/utils"
)

type ZodCodegen struct{}

func New() *ZodCodegen {
	return &ZodCodegen{}
}

func (generator *ZodCodegen) Generate(root domain.Type) string {
	code := &strings.Builder{}
	code.WriteString("import { z } from \"zod\";")
	code.WriteRune('\n')
	code.WriteRune('\n')
	visited := map[string]string{}
	return generator.dfs(root, visited, code)
}

func (generator *ZodCodegen) dfs(node domain.Type, visited map[string]string, code *strings.Builder) string {
	id := node.Canonical()
	if val, ok := visited[id]; ok {
		return val
	}

	switch node := node.(type) {
	case *domain.ObjectType:
		builder := new_schema_builder()
		builder.with_name("Root")
		for key, field := range node.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}
		code.WriteString(builder.build())

		return code.String()

	case *domain.NamedType:
		type_name := utils.Capitalize(node.Namespace)

		builder := new_schema_builder()
		builder.with_name(type_name)

		for key, field := range node.Identity.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}

		code.WriteString(builder.build())
		code.WriteRune('\n')
		visited[id] = builder.schema_name()

		return builder.schema_name()

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

		union_type := fmt.Sprintf("z.union([%s])", strings.Join(types, ", "))
		visited[id] = union_type
		return union_type
	case *domain.ArrayType:
		field_type := generator.dfs(node.Element, visited, code)
		type_def := fmt.Sprintf("z.array(%s)", field_type)
		visited[id] = type_def
		return type_def
	case *domain.IntType, *domain.FloatType, domain.IntType, domain.FloatType:
		visited[id] = "z.number()"
		return "z.number()"
	case *domain.NullType, domain.NullType:
		visited[id] = "z.null()"
		return "z.null()"
	case *domain.UnknownType, domain.UnknownType:
		visited[id] = "z.unknown()"
		return "z.unknown()"
	case *domain.DateType, domain.DateType:
		visited[id] = "z.date()"
		return "z.date()"
	case *domain.StringType, domain.StringType:
		visited[id] = "z.string()"
		return "z.string()"
	case *domain.BooleanType, domain.BooleanType:
		visited[id] = "z.boolean()"
		return "z.boolean()"
	default:
		return fmt.Sprintf("Unhandled type '%s'", node.Canonical())
	}

}
