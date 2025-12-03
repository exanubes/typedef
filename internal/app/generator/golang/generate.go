package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/transformer"
)

type GolangCodegen struct {
	transformer transformer.Transformer
}

func New(transformer transformer.Transformer) *GolangCodegen {
	return &GolangCodegen{transformer: transformer}
}

func (generator *GolangCodegen) Generate(tree *ast.Program) string {
	typedef := generator.transformer.Transform(tree)
	var builder strings.Builder
	type_map := map[string]string{}
	// TODO: deterministic ordering of properties
	for index, node := range typedef {
		type_map[node.ID] = capitalize(node.Name)
		builder.WriteString(fmt.Sprintf("type %s struct {", capitalize(node.Name)))
		builder.WriteRune('\n')
		for _, field := range node.Fields {
			builder.WriteString(fmt.Sprintf("  %s %s", capitalize(field.Name), generator.parse_type(field.TypeID, type_map)))
			builder.WriteRune('\n')
		}
		builder.WriteString("}")
		if index != len(typedef) {
			builder.WriteRune('\n')
			builder.WriteRune('\n')
		}
	}

	return builder.String()
}

func (generator *GolangCodegen) parse_type(value string, types map[string]string) string {
	if strings.HasPrefix(value, "named") {
		return types[value]
	}

	return value
}
func capitalize(input string) string {
	if input == "id" {
		return "ID"
	}

	return strings.ToUpper(input[:1]) + input[1:]
}
