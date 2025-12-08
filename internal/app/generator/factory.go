package generator

import (
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/generator/typescript"
)

type CodegenFactory struct{}

func (factory CodegenFactory) Create(format Format) CodeGenerator {
	switch format {
	case GOLANG:
		return golang.New()
	case TYPESCRIPT:
		return typescript.New()
	}

	return nil
}

type Format int

const (
	GOLANG = iota
	TYPESCRIPT
)
