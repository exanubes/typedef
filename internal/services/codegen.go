package services

import (
	"fmt"

	"github.com/exanubes/typedef/internal/app/generator"
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer"
	"github.com/exanubes/typedef/internal/app/parser"
	"github.com/exanubes/typedef/internal/app/transformer"
	"github.com/exanubes/typedef/internal/domain"
)

type CodegenService struct {
	lexer       lexer.Factory
	parser      parser.Factory
	graph       graph.TypeGraph
	transformer transformer.Transformer
	codegen     generator.CodeGenerator
}

func NewCodegenService(
	lexer lexer.Factory,
	parser parser.Factory,
	graph graph.TypeGraph,
	transformer transformer.Transformer,
	codegen generator.CodeGenerator,
) *CodegenService {
	return &CodegenService{
		lexer:       lexer,
		parser:      parser,
		graph:       graph,
		transformer: transformer,
		codegen:     codegen,
	}
}

func (service *CodegenService) Execute(options domain.CodegenOptions) (string, error) {
	lexer := service.lexer.Create(options.InputType, options.Input)
	if lexer == nil {
		return "", fmt.Errorf("%s format is not supported", options.InputType)
	}
	parser := service.parser.Create(options.InputType, lexer)
	if parser == nil {
		return "", fmt.Errorf("%s format is not supported", options.InputType)
	}
	ast := parser.Parse()
	graph_root := service.graph.Generate(ast)
	v2 := golang.NewV2()
	return v2.Generate("Root", graph_root), nil
}
