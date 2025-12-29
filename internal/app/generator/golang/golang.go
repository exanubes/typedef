package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/utils"
)

type GolangCodegen struct {
	id_provider *utils.IdProvider
}

func New() *GolangCodegen {
	return &GolangCodegen{
		id_provider: &utils.IdProvider{},
	}
}

func (generator *GolangCodegen) Generate(root domain.Type) string {
	code := &strings.Builder{}
	visited := map[string]string{}
	generator.id_provider.Reset()
	return generator.dfs(root, visited, code)
}

func (generator *GolangCodegen) dfs(node domain.Type, visited map[string]string, code *strings.Builder) string {
	id := node.Canonical()
	if val, ok := visited[id]; ok {
		return val
	}

	switch node := node.(type) {
	case *domain.ObjectType:
		builder := new_struct_builder()
		builder.with_name("Root")
		for key, field := range node.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}
		code.WriteString(builder.build())

		return code.String()

	case *domain.NamedType:
		struct_name := utils.Capitalize(node.Namespace)

		builder := new_struct_builder()
		builder.with_name(struct_name)

		for key, field := range node.Identity.Fields {
			field_type := generator.dfs(field, visited, code)
			builder.with_field(key, field_type)
		}

		code.WriteString(builder.build())
		code.WriteRune('\n')
		visited[id] = struct_name

		return struct_name

	case *domain.UnionType:
		struct_name := fmt.Sprintf("UnionType_%d", generator.id_provider.Next())

		builder := new_struct_builder()
		builder.with_name(struct_name)
		types := []string{}
		for _, typ := range node.OneOf {
			types = append(types, generator.dfs(typ, visited, code))
		}
		types = utils.SortFields(types)

		for index, union_type := range types {
			builder.with_field(utils.Letter(index), union_type)
		}

		visited[id] = struct_name
		code.WriteString(builder.build())
		code.WriteRune('\n')

		return struct_name

	case *domain.ArrayType:
		field_type := generator.dfs(node.Element, visited, code)
		type_def := fmt.Sprintf("[]%s", field_type)
		visited[id] = type_def
		return type_def
	case *domain.FloatType, domain.FloatType:
		visited[id] = "float64"
		return "float64"
	case *domain.BooleanType, domain.BooleanType:
		visited[id] = "bool"
		return "bool"
	case *domain.NullType, domain.NullType:
		visited[id] = "nil"
		return "nil"
	case *domain.UnknownType, domain.UnknownType:
		visited[id] = "any"
		return "any"
	case *domain.DateType, domain.DateType:
		visited[id] = "time.Time"
		return "time.Time"
	case *domain.UuidType, domain.UuidType:
		visited[id] = "string"
		// TODO: make configurable to use external type e.g., github.com/google/uuid
		return "string"
	default:
		visited[id] = node.Canonical()
		return node.Canonical()
	}

}
