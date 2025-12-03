package transformer

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
	}
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()

	graph := graph.New(dedup.New())
	transformer := New(graph)
	result := transformer.Transform(program)

	if len(result) != 2 {
		test.Fatalf("Expected 2 type definitions, received %d", len(result))
	}

	if result[0].ID != "named(user#b7c1f9c4)" {
		test.Fatalf("Expected result[0].ID to be \"name(user#b7c1f9c4)\", received: %s", result[0].ID)
	}

	if result[1].ID != "root" {
		test.Fatalf("Expected result[1].ID to be \"root\", received: %s", result[1].ID)
	}
}
