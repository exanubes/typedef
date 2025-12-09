package golang

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/utils"
	"golang.org/x/exp/maps"
)

type struct_builder struct {
	name   string
	fields map[string]string
}

var indent = "  "

func new_struct_builder() *struct_builder {
	return &struct_builder{fields: make(map[string]string)}
}

func (builder *struct_builder) with_name(struct_name string) *struct_builder {
	builder.name = utils.Capitalize(struct_name)
	return builder
}

func (builder *struct_builder) with_field(field_name, field_type string) *struct_builder {
	builder.fields[field_name] = indent + fmt.Sprintf("%s %s", utils.Capitalize(field_name), field_type)
	return builder
}

func (builder *struct_builder) build() string {
	var code strings.Builder
	fields := utils.SortFields(maps.Keys(builder.fields))
	code.WriteString(fmt.Sprintf("type %s struct {", builder.name))
	code.WriteRune('\n')
	for _, field := range fields {
		value := builder.fields[field]
		code.WriteString(value)
		code.WriteRune('\n')
	}
	code.WriteString("}")
	code.WriteRune('\n')

	return code.String()
}
