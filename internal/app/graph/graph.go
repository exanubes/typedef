package graph

import (
	"log"
	"sort"

	"golang.org/x/exp/maps"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/domain"
)

type Graph struct {
	type_pool dedup.Engine
}

var _ TypeGraph = (*Graph)(nil)

func New(type_pool dedup.Engine) *Graph {
	return &Graph{type_pool: type_pool}
}

func (graph *Graph) Generate(tree *ast.Program) *domain.ObjectType {
	switch node := tree.Value.(type) {
	case *ast.ObjectNode:
		return graph.parse_object(node)
	default:
		log.Fatalf("Invalid node %+v", node)
	}
	return nil
}

func (graph *Graph) parse_object(node *ast.ObjectNode) *domain.ObjectType {
	result := domain.ObjectType{Fields: make(map[string]domain.Type)}
	for _, pair := range node.Pairs {
		result.Fields[pair.Key.Value] = graph.parse_value(pair.Key.Value, pair.Value)
	}
	return &result
}

func (graph *Graph) parse_value(property string, node ast.Node) domain.Type {
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
		object_type := graph.parse_object(node)
		named_type := graph.type_pool.Get(object_type)
		if named_type != nil {
			return named_type
		}

		named_type = &domain.NamedType{Identity: object_type, Namespace: property}
		graph.type_pool.Add(named_type)

		return named_type
	case *ast.ArrayNode:
		types_map := map[string]domain.Type{}
		if len(node.Elements) == 0 {
			return &domain.ArrayType{Element: domain.UnknownType{}}
		}

		for _, element := range node.Elements {
			value := graph.parse_value(property, element)
			types_map[value.Canonical()] = value
		}

		types := maps.Values(types_map)
		if len(types) > 1 {
			// Sort types deterministically by their canonical representation
			sort.Slice(types, func(i, j int) bool {
				return types[i].Canonical() < types[j].Canonical()
			})
			result = &domain.ArrayType{Element: &domain.UnionType{OneOf: types}}
		} else {
			result = &domain.ArrayType{Element: types[0]}
		}
	}
	return result
}
