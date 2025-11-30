package graph

import (
	"fmt"
	"testing"

	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
)

func TestGraphTypeNodes(test *testing.T) {
	input := `
	{
	"bool": true,
	"int": 69420,
	"float": 69.420,
	"varchar": "hello world"
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()
	fmt.Printf("PROGRAM: %+v\n", program)
	graph := Generate(program)

	testcases := []struct {
		property     string
		expectedType string
	}{
		{"bool", "boolean"},
		{"int", "int"},
		{"float", "float"},
		{"varchar", "string"},
	}

	for index, tc := range testcases {
		if graph.Fields[tc.property].Name() != tc.expectedType {
			test.Fatalf("tests[%d] - wrong graph type node, expected=%q, received %q", index, tc.expectedType, graph.Fields[tc.property].Name())
		}
	}
}

func TestGraphArrayTypeNodes(test *testing.T) {
	input := `
	{
	"list": [1,2,3],
	"union": [1, "2", true]
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()
	graph := Generate(program)

	testcases := []struct {
		property     string
		expectedType string
	}{
		{"list", "array"},
		{"union", "array"},
	}

	for index, tc := range testcases {
		if graph.Fields[tc.property].Name() != tc.expectedType {
			test.Fatalf("tests[%d] - wrong graph type node, expected=%q, received %q", index, tc.expectedType, graph.Fields[tc.property].Name())
		}
	}
}
