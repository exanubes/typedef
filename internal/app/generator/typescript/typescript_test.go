package typescript

import (
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
)

func TestNamedTypes(test *testing.T) {
	input := `{
	"id": 1,
	"title": "Harry Potter",
	"user":{"id": 1,
	"name": "John"
	},
	"author": {
	"id": 2,
	"name": "Tom"
	},
	"numbers": [1,2,3],
	"mixed": [1, "2", false, {"id": 3, "name": "Simon"}],
	"float": 69.420,
	"cool": true
	}
	`

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
  cool: boolean;
  float: number;
  mixed: (User | boolean | number | string)[];
  numbers: number[];
  title: string;
  user: User;
}
`

	if result != expected {
		test.Fatalf("Expected \n%s, received \n%s", expected, result)
	}

}
