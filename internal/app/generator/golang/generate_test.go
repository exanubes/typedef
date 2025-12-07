package golang

import (
	"fmt"
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
	"github.com/exanubes/typedef/internal/app/transformer"
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
	"arr": [1, "2", true]
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	graph := graph.New(dedup.New())
	transformer := transformer.New()
	codegen := New()
	fmt.Printf("IR: %+v", transformer.Transform(graph.Generate(parser.Parse())))
	result := codegen.Generate(transformer.Transform(graph.Generate(parser.Parse())))
	expected := `type User struct {
  ID int
  Name string
}

type Root struct {
  ID int
  Author User
  Title string
  User User
}`

	if result != expected {
		test.Fatalf("Expected \n%s, received \n%s", expected, result)
	}

}
