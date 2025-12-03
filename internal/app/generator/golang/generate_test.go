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
	}
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()

	graph := graph.New(dedup.New())
	transformer := transformer.New(graph)
	codegen := New(transformer)
	result := codegen.Generate(program)
	expected := `type User struct {
  ID int
  Name string
}

type Root struct {
  User User
  Author User
  ID int
  Title string
}`
	if result != expected {
		test.Fatalf("Expected %s, received %s", expected, result)
	}

}
