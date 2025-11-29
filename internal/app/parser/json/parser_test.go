package parser

import (
	"testing"

	"github.com/exanubes/typedef/internal/app/ast"
	"github.com/exanubes/typedef/internal/app/lexer/json"
)

func TestNextToken(test *testing.T) {
	input := `
	[{
	"property": true,
	"prop1": null
	}]`

	lexer := json.New(input)
	parser := New(lexer)

	program := parser.Parse()

	print("object", program.Value.String())

	array, ok := program.Value.(*ast.ArrayNode)

	if !ok {
		test.Fatalf("expect array, received %+v", array)
	}

	if len(array.Elements) != 1 {
		test.Fatalf("expected array to have 1 element, received %d", len(array.Elements))
	}

	object, ok := array.Elements[0].(*ast.ObjectNode)

	if !ok {
		test.Fatalf("expected array[0] to be an object, received %+v", object)
	}

	if len(object.Pairs) != 2 {
		test.Fatalf("expected object to have 2 key/value pairs, received %d", len(object.Pairs))
	}

	if object.Pairs[0].Key.String() != "property" {
		test.Fatalf("expected property, received %s", object.Pairs[0].Key.String())
	}

	if object.Pairs[0].Value.String() != "true" {
		test.Fatalf("expected true, received %s", object.Pairs[0].Value.String())
	}

	if object.Pairs[1].Key.String() != "prop1" {
		test.Fatalf("expected property, received %s", object.Pairs[0].Key.String())
	}

	if object.Pairs[1].Value.String() != "null" {
		test.Fatalf("expected true, received %s", object.Pairs[0].Value.String())
	}

}
