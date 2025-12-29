package jsdoc

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/utils"
	"github.com/exanubes/typedef/internal/utils/maps"
)

type typedef_builder struct {
	name   string
	fields map[string]string
}

var indent = " * "

func new_typedef_builder() *typedef_builder {
	return &typedef_builder{fields: make(map[string]string)}
}

func (builder *typedef_builder) with_name(type_name string) *typedef_builder {
	builder.name = utils.Capitalize(type_name)
	return builder
}

func (builder *typedef_builder) with_field(field_name, field_type string) *typedef_builder {
	builder.fields[field_name] = indent + fmt.Sprintf("@property { %s } %s", field_type, field_name)
	return builder
}

func (builder *typedef_builder) build() string {
	fields := utils.SortFields(maps.Keys(builder.fields))
	var code strings.Builder
	code.WriteString("/**")
	code.WriteRune('\n')
	code.WriteString(indent + fmt.Sprintf("@typedef  { Object } %s", builder.name))
	code.WriteRune('\n')

	for _, field := range fields {
		value := builder.fields[field]
		code.WriteString(value)
		code.WriteRune('\n')
	}

	code.WriteString(" */")
	code.WriteRune('\n')

	return code.String()
}
