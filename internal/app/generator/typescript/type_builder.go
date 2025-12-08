package typescript

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/utils"
	"golang.org/x/exp/maps"
)

type type_builder struct {
	name   string
	fields map[string]string
}

var indent = "  "

func new_type_builder() *type_builder {
	return &type_builder{fields: make(map[string]string)}
}

func (builder *type_builder) with_name(type_name string) *type_builder {
	builder.name = utils.Capitalize(type_name)
	return builder
}

func (builder *type_builder) with_field(field_name, field_type string) *type_builder {
	builder.fields[field_name] = indent + fmt.Sprintf("%s: %s;", field_name, field_type)
	return builder
}

func (builder *type_builder) build() string {
	fields := utils.SortFields(maps.Keys(builder.fields))
	var code strings.Builder

	code.WriteString(fmt.Sprintf("type %s = {", builder.name))
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
