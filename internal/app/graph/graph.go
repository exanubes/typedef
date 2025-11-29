package graph

import (
	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/domain"
	"log"
)

func Generate(tree *ast.Program) domain.ObjectType {
	var result domain.ObjectType

	switch node := tree.Value.(type) {
	case ast.ObjectNode:
		for _, pair := range node.Pairs {
			var vt domain.Type
			switch val := pair.Value.(type) {
			case ast.StringNode:
				// TODO: date type
				vt = domain.StringType{}
			case ast.NumberNode:
				switch val.Kind {
				case ast.INTEGER:
					vt = domain.IntType{}
				case ast.FLOAT:
					vt = domain.FloatType{}
				}
			case ast.BooleanNode:
				vt = domain.BooleanType{}
			}
			result.Fields[pair.Key.Value] = vt
		}
	default:
		log.Fatalf("Invalid node %+v", node)
	}
	return result
}
