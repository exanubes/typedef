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
	// TODO: deterministic ordering of struct properties
	// TODO: union types
	case *domain.ObjectType:
		var codegen strings.Builder
		codegen.WriteString("type Root struct {")
		codegen.WriteRune('\n')
		for key, field := range node.Fields {
			field_type := generator.dfs(field, visited, code)
			codegen.WriteString(indent + fmt.Sprintf("%s %s", capitalize(key), field_type))
			codegen.WriteRune('\n')
		}
		codegen.WriteString("}")
		codegen.WriteRune('\n')
		code.WriteString(codegen.String())
		return code.String()
	case *domain.NamedType:
		var codegen strings.Builder
		struct_name := capitalize(node.Namespace)
		codegen.WriteString(fmt.Sprintf("type %s struct {", struct_name))
		codegen.WriteRune('\n')
		for key, field := range node.Identity.Fields {
			field_type := generator.dfs(field, visited, code)
			codegen.WriteString(indent + fmt.Sprintf("%s %s", capitalize(key), field_type))
			codegen.WriteRune('\n')
		}
		codegen.WriteString("}")
		codegen.WriteRune('\n')
		codegen.WriteRune('\n')
		code.WriteString(codegen.String())
		visited[id] = struct_name
		return struct_name
	// FIX: Inconsistency with pointer/value receivers
	case domain.ArrayType:
		field_type := generator.dfs(node.Element, visited, code)
		type_def := fmt.Sprintf("[]%s", field_type)
		visited[id] = type_def
		return type_def
	default:
		visited[id] = node.Canonical()
		return node.Canonical()
	}
}
