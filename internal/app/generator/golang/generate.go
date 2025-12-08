package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
)

type GolangCodegen struct{}

func New() *GolangCodegen {
	return &GolangCodegen{}
}

var indent = "  "

func (generator *GolangCodegen) Generate(root domain.Type) string {
	code := &strings.Builder{}
	visited := map[string]string{}
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
		struct_name := capitalize(node.Namespace)

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
		struct_name := fmt.Sprintf("UnionType_%s", random_string(10))

		builder := new_struct_builder()
		builder.with_name(struct_name)

		for index, typ := range node.OneOf {
			union_type := generator.dfs(typ, visited, code)
			builder.with_field(string(alpha[index]), union_type)
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
	default:
		visited[id] = node.Canonical()
		return node.Canonical()
	}

}
