package generator

import (
	"github.com/exanubes/typedef/internal/app/transformer"
)

type CodeGenerator interface {
	Generate([]transformer.TypeDef) string
}
