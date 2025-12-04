package transformer

import (
	"github.com/exanubes/typedef/internal/domain"
	"golang.org/x/exp/maps"
)

type IntermediateRepresentationTransformer struct {
}

var _ Transformer = (*IntermediateRepresentationTransformer)(nil)

func New() *IntermediateRepresentationTransformer {
	return &IntermediateRepresentationTransformer{}
}

func (transformer *IntermediateRepresentationTransformer) Transform(root *domain.ObjectType) []TypeDef {

	visited := map[string]bool{}
	typedefs := map[string]TypeDef{}
	transformer.dfs(root, visited, typedefs)

	return maps.Values(typedefs)
}

func (transformer *IntermediateRepresentationTransformer) dfs(node domain.Type, visited map[string]bool, defs map[string]TypeDef) {
	id := node.Canonical()
	if visited[id] {
		return
	}
	visited[id] = true

	switch node := node.(type) {
	case *domain.ObjectType:
		for _, field := range node.Fields {
			transformer.dfs(field, visited, defs)
		}
		defs[id] = objectToTypeDef("root", "root", node.Fields)
	case *domain.NamedType:
		for _, field := range node.Identity.Fields {
			transformer.dfs(field, visited, defs)
		}

		defs[id] = objectToTypeDef(node.Canonical(), node.Namespace, node.Identity.Fields)
	case domain.ArrayType:
		transformer.dfs(node.Element, visited, defs)
	case domain.UnionType:
		for _, field := range node.OneOf {
			transformer.dfs(field, visited, defs)
		}
		defs[id] = unionToTypeDef(node)
	default:
		return
	}
}
