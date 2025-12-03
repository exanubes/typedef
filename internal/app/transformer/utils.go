package transformer

import "github.com/exanubes/typedef/internal/domain"

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
	return TypeDef{
		ID:     id,
		Name:   name,
		Fields: fields,
	}
}

func unionToTypeDef(field domain.UnionType) TypeDef {
	fields := make([]FieldDef, len(field.OneOf))
	for index, field := range field.OneOf {
		fields[index] = FieldDef{
			Name:       "",
			ParsedName: "",
			TypeID:     field.Canonical(),
		}
	}
	return TypeDef{
		ID:     field.Canonical(),
		Name:   field.Canonical(),
		Fields: fields,
	}
}
