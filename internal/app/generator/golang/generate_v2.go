package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
)

type GolangCodegenV2 struct{}

func NewV2() *GolangCodegenV2 {
	return &GolangCodegenV2{}
}

var indent = "  "

func (generator *GolangCodegenV2) Generate(type_name string, root *domain.ObjectType) string {
	code := &strings.Builder{}
	visited := map[string]string{}
	return generator.dfs(root, visited, code)
}
func (generator *GolangCodegenV2) dfs(node domain.Type, visited map[string]string, code *strings.Builder) string {
	id := node.Canonical()
	if val, ok := visited[id]; ok {
		return val
	}

	switch node := node.(type) {
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
