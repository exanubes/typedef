package generator

import (
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/generator/jsdoc"
	"github.com/exanubes/typedef/internal/app/generator/typescript"
	"github.com/exanubes/typedef/internal/app/generator/zod"
	"github.com/exanubes/typedef/internal/utils"
)

type CodegenFactory struct{}

func (factory CodegenFactory) Create(format Format) CodeGenerator {
	switch format {
	case GOLANG:
		return golang.New(utils.RandomString)
	case TYPESCRIPT:
		return typescript.New()
	case ZOD:
		return zod.New()
	case JSDOC:
		return jsdoc.New()
	}

	return nil
}

type Format int

const (
	GOLANG = iota
	TYPESCRIPT
	ZOD
	JSDOC
)
