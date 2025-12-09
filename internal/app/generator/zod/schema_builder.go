package zod

import (
	"fmt"
	"strings"

	"github.com/exanubes/typedef/internal/utils"
	"golang.org/x/exp/maps"
)

type schema_builder struct {
	name   string
	fields map[string]string
}

var indent = "  "

func new_schema_builder() *schema_builder {
	return &schema_builder{fields: make(map[string]string)}
}

func (builder *schema_builder) with_name(type_name string) *schema_builder {
	builder.name = utils.Capitalize(type_name)
	return builder
}

func (builder *schema_builder) with_field(field_name, field_type string) *schema_builder {
	builder.fields[field_name] = indent + fmt.Sprintf("%s: %s,", field_name, field_type)
	return builder
}

func (builder *schema_builder) schema_name() string {
	return fmt.Sprintf("%sSchema", builder.name)
}

func (builder *schema_builder) build() string {
	fields := utils.SortFields(maps.Keys(builder.fields))
	var code strings.Builder

	code.WriteString(fmt.Sprintf("const %s = z.object({", builder.schema_name()))
	code.WriteRune('\n')

	for _, field := range fields {
		value := builder.fields[field]
		code.WriteString(value)
		code.WriteRune('\n')
	}

	code.WriteString("});")
	code.WriteRune('\n')
	code.WriteString(fmt.Sprintf("type %s = z.infer<typeof %s>;", builder.name, builder.schema_name()))
	code.WriteRune('\n')

	return code.String()
}
