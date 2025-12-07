package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/app/transformer"
)

type GolangCodegen struct{}

func New() *GolangCodegen {
	return &GolangCodegen{}
}

func (generator *GolangCodegen) Generate(typedef []transformer.TypeDef) string {
	var builder strings.Builder
	type_map := map[string]transformer.TypeDef{}
	for _, node := range typedef {
		type_map[node.ID] = node
	}

	fmt.Printf("INPUT: %+v\n\n", typedef)
	for index, node := range typedef {
		if node.Kind == transformer.KindArray {
			continue
		}

		builder.WriteString(fmt.Sprintf("type %s struct {", capitalize(node.Name)))
		builder.WriteRune('\n')
		for _, field := range node.Fields {
			builder.WriteString(
				fmt.Sprintf("  %s %s",
					capitalize(field.Name),
					parse_type(field.TypeID, type_map),
				))
			builder.WriteRune('\n')
		}
		builder.WriteString("}")
		if index != len(typedef)-1 {
			builder.WriteRune('\n')
			builder.WriteRune('\n')
		}
	}

	return builder.String()
}

func parse_type(value string, types map[string]transformer.TypeDef) string {
	if strings.HasPrefix(value, "named") {
		return capitalize(types[value].Name)
	}

	if strings.HasPrefix(value, "array") {
		return fmt.Sprintf("[]%s", parse_type(types[value].ElementType, types))
	}

	return value
}
func capitalize(input string) string {
	if input == "id" {
		return "ID"
	}

	return strings.ToUpper(input[:1]) + input[1:]
}
