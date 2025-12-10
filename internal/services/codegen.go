package services

import (
	"fmt"

	"github.com/exanubes/typedef/internal/app/generator"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer"
	"github.com/exanubes/typedef/internal/app/parser"
	"github.com/exanubes/typedef/internal/domain"
)

type CodegenService struct {
	lexer   lexer.Factory
	parser  parser.Factory
	graph   graph.TypeGraph
	codegen generator.Factory
}

func NewCodegenService(
	lexer lexer.Factory,
	parser parser.Factory,
	graph graph.TypeGraph,
	codegen generator.Factory,
) *CodegenService {
	return &CodegenService{
		lexer:   lexer,
		parser:  parser,
		graph:   graph,
		codegen: codegen,
	}
}

func (service *CodegenService) Execute(options domain.CodegenOptions) (string, error) {
	lexer := service.lexer.Create(options.InputType, options.Input)
	if lexer == nil {
		return "", fmt.Errorf("%s lexer is not supported", options.InputType)
	}

	parser := service.parser.Create(options.InputType, lexer)
	if parser == nil {
		return "", fmt.Errorf("%s parser is not supported", options.InputType)
	}

	ast := parser.Parse()
	graph_root := service.graph.Generate(ast)

	codegen := service.codegen.Create(options.Format)
	if codegen == nil {
		return "", fmt.Errorf("%s format is not supported", options.Format)
	}

	return codegen.Generate(graph_root), nil
}
