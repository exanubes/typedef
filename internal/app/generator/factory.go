package generator

import (
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/generator/jsdoc"
	"github.com/exanubes/typedef/internal/app/generator/typescript"
	"github.com/exanubes/typedef/internal/app/generator/zod"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/utils"
)

type CodegenFactory struct{}

func (factory CodegenFactory) Create(format domain.Format) CodeGenerator {
	switch format {
	case domain.GOLANG:
		return golang.New(utils.RandomString)
	case domain.TYPESCRIPT:
		return typescript.New()
	case domain.ZOD:
		return zod.New()
	case domain.JSDOC:
		return jsdoc.New()
	}

	return nil
}
