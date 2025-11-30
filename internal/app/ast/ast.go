package ast

import "fmt"

type Node interface {
	Literal() string
}

type NumberKind string

const (
	FLOAT   = "float"
	INTEGER = "integer"
)

type ObjectNode struct {
	Pairs []*PairNode
}

func (node ObjectNode) Literal() string {
	pairs := []PairNode{}

	for _, pair := range node.Pairs {
		pairs = append(pairs, *pair)
	}

	return fmt.Sprintf("%+v", pairs)
}

type PairNode struct {
	Key   *StringNode
	Value Node
}

type ArrayNode struct {
	Elements []Node
}

func (node ArrayNode) Literal() string {
	return "[array Array]"
}

type StringNode struct {
	Value string
}

func (node StringNode) Literal() string {
	return node.Value
}

type BooleanNode struct {
	Value string
}

func (node BooleanNode) Literal() string {
	return node.Value
}

type NumberNode struct {
	Value string
	Kind  NumberKind
}

func (node NumberNode) Literal() string {
	return node.Value
}

type NullNode struct{}

func (node NullNode) Literal() string {
	return "null"
}

type Program struct {
	Value Node
}

func (program *Program) Literal() string {
	if program.Value == nil {
		return ""
	}

	return program.Value.Literal()
}
