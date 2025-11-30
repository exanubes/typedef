package graph

import (
	"golang.org/x/exp/maps"
	"log"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/domain"
)

func Generate(tree *ast.Program) domain.ObjectType {
	result := domain.ObjectType{Fields: make(map[string]domain.Type)}
	switch node := tree.Value.(type) {
	case *ast.ObjectNode:
		for _, pair := range node.Pairs {
			result.Fields[pair.Key.Value] = parse_value(pair.Value)
		}
	default:
		log.Fatalf("@@ Invalid node %+v", node)
	}
	return result
}

func parse_value(node ast.Node) domain.Type {
	var result domain.Type
	switch node := node.(type) {
	case *ast.StringNode:
		// TODO: date type
		result = domain.StringType{}
	case *ast.NumberNode:
		switch node.Kind {
		case ast.INTEGER:
			result = domain.IntType{}
		case ast.FLOAT:
			result = domain.FloatType{}
		}
	case *ast.BooleanNode:
		result = domain.BooleanType{}
	case *ast.ObjectNode:
		//TODO: build a named node and compare with existing named nodes and dedup
	case *ast.ArrayNode:
		types_map := map[string]domain.Type{}
		for _, element := range node.Elements {
			value := parse_value(element)
			types_map[value.Name()] = value
		}
		types := maps.Values(types_map)
		if len(types) > 1 {
			result = domain.ArrayType{Element: domain.UnionType{OneOf: types}}
		} else {
			result = domain.ArrayType{Element: types[0]}
		}
	}
	return result
}
