package ast

type Node interface {
	String() string
}

type NumberKind string

const (
	FLOAT   = "float"
	INTEGER = "integer"
)

type ObjectNode struct {
	Pairs []*PairNode
}

func (node ObjectNode) String() string {
	return "{}"
}

type PairNode struct {
	Key   *StringNode
	Value Node
}

type ArrayNode struct {
	Elements []Node
}

func (node ArrayNode) String() string {
	return ""
}

type StringNode struct {
	Value string
}

func (node StringNode) String() string {
	return node.Value
}

type BooleanNode struct {
	Value string
}

func (node BooleanNode) String() string {
	return node.Value
}

type NumberNode struct {
	Value string
	Kind  NumberKind
}

func (node NumberNode) String() string {
	return node.Value
}

type NullNode struct{}

func (node NullNode) String() string {
	return "null"
}

type Program struct {
	Value Node
}

func (program *Program) String() string {
	if program.Value == nil {
		return ""
	}

	return program.Value.String()
}
