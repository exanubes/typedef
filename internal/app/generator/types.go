package generator

import (
	"github.com/exanubes/typedef/internal/domain"
)

type CodeGenerator interface {
	Generate(domain.Type) string
}
