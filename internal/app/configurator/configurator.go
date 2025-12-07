package configurator

import (
	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer"
	"github.com/exanubes/typedef/internal/app/parser"
	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/infrastructure/targets"
	"github.com/exanubes/typedef/internal/services"
)

func New() (domain.CodegenService, domain.OutputTarget) {
	type_pool := dedup.New()
	return services.NewCodegenService(
			lexer.LexerFactory{},
			parser.ParserFactory{},
			graph.New(type_pool),
			golang.New(),
		),
		targets.Create("json")
}
