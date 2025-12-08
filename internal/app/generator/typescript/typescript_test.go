package typescript

import (
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
)

func TestNamedTypes(test *testing.T) {
	input := `
	{
	"id": 1,
	"title": "Harry Potter",
	"user": {
	"id": 1,
	"name": "John"
	},
	"author": {
	"id": 2,
	"name": "Tom"
	},
	"mixed": [1, "two", true, {"id": 69420, "name": "Jane"}],
	"unary": [6,9,4,2,0]
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	graph := graph.New(dedup.New())
	codegen := NewTypescriptCodegen()
	result := codegen.Generate(graph.Generate(parser.Parse()))
	expected := `type User = {
  id: number;
  name: string;
}

type Root = {
  id: number;
  author: User;
  mixed: (User | boolean | number | string)[];
  title: string;
  unary: number[];
  user: User;
}
`

	if result != expected {
		test.Fatalf("Expected \n%s, received \n%s", expected, result)
	}

}
