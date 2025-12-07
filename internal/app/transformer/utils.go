package transformer

import (
	"slices"
	"strings"

	"github.com/exanubes/typedef/internal/domain"
)

func objectToTypeDef(id, name string, properties map[string]domain.Type) TypeDef {
	fields := make([]FieldDef, len(properties))
	index := 0
	for key, field := range properties {
		fields[index] = FieldDef{
			Name:       key,
			ParsedName: key,
			TypeID:     field.Canonical(),
		}

		index += 1
	}
	slices.SortStableFunc(fields, func(a, b FieldDef) int {
		if strings.ToLower(a.Name) == "id" {
			return -1
		}

		if strings.ToLower(b.Name) == "id" {
			return 1
		}

		if a.Name > b.Name {
			return 1
		}

		if a.Name < b.Name {
			return -1
		}

		return 0
	})
	return TypeDef{
		ID:     id,
		Name:   name,
		Kind:   KindObject,
		Fields: fields,
	}
}

var letters = []string{"A", "B", "C", "D", "E", "F"}

func unionToTypeDef(field domain.UnionType) TypeDef {
	fields := make([]FieldDef, len(field.OneOf))
	for index, field := range field.OneOf {
		fields[index] = FieldDef{
			Name:       letters[index],
			ParsedName: "",
			TypeID:     field.Canonical(),
		}
	}
	return TypeDef{
		ID:     field.Canonical(),
		Name:   field.Canonical(),
		Fields: fields,
		Kind:   KindUnion,
	}
}

func arrayToTypeDef(field domain.ArrayType) TypeDef {
	return TypeDef{
		ID:          field.Canonical(),
		Name:        field.Canonical(),
		Kind:        KindArray,
		ElementType: field.Element.Canonical(),
	}
}
