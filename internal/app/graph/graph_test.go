package graph

import (
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
	"github.com/exanubes/typedef/internal/domain"
)

func TestCanonicalSerialization(test *testing.T) {
	input := `
	{
	"bool": true,
	"int": 69420,
	"float": 69.420,
	"varchar": "hello world",
	"simple": [1,2,3],
	"mixed": [69.420, 69, 420, "hello", "there", true, false]
	}`
	expected := "object{bool:boolean,float:float,int:int,mixed:array<union<boolean|float|int|string>>,simple:array<int>,varchar:string}"
	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()

	engine := New(dedup.New())
	graph := engine.Generate(program)

	if result := graph.Canonical(); result != expected {
		test.Fatalf("Expected %s, received: %s", expected, result)
	}
}

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
	engine := New(dedup.New())
	graph := engine.Generate(program)
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
	"union": [1, "2", true],
	"numbers": [69420, 69.420],
	"empty": []
	}`

	lexer := json.New(input)
	parser := parser.New(lexer)

	program := parser.Parse()
	engine := New(dedup.New())
	graph := engine.Generate(program)

	testcases := []struct {
		property     string
		expectedType string
	}{
		{"list", "array"},
		{"union", "array"},
		{"numbers", "array"},
		{"empty", "array"},
	}

	for index, tc := range testcases {
		if graph.Fields[tc.property].Name() != tc.expectedType {
			test.Fatalf("tests[%d] - wrong graph type node, expected=%q, received %q", index, tc.expectedType, graph.Fields[tc.property].Name())
		}
	}

	unary_array := graph.Fields["list"].(domain.ArrayType)

	if unary_array.Element.Name() != (domain.IntType{}).Name() {
		test.Fatalf("Wrong array element type expected %s, received %s", (domain.IntType{}).Name(), unary_array.Element.Name())
	}

	union_array := graph.Fields["union"].(domain.ArrayType)

	if union_array.Element.Name() != (domain.UnionType{}).Name() {
		test.Fatalf("Wrong array element type expected %s, received %s", (domain.UnionType{}).Name(), union_array.Element.Name())
	}

	union := union_array.Element.(domain.UnionType)

	if union.OneOf[0].Name() != (domain.IntType{}).Name() {
		test.Fatalf("Wrong union[0] type expected %s, received %s", (domain.IntType{}).Name(), union_array.Element.Name())
	}
	if union.OneOf[1].Name() != (domain.StringType{}).Name() {
		test.Fatalf("Wrong union[1] type expected %s, received %s", (domain.StringType{}).Name(), union_array.Element.Name())
	}
	if union.OneOf[2].Name() != (domain.BooleanType{}).Name() {
		test.Fatalf("Wrong union[2] type expected %s, received %s", (domain.BooleanType{}).Name(), union_array.Element.Name())
	}

	numbers_array := graph.Fields["numbers"].(domain.ArrayType)

	if numbers_array.Element.Name() != (domain.UnionType{}).Name() {
		test.Fatalf("Wrong array element type expected %s, received %s", (domain.UnionType{}).Name(), numbers_array.Element.Name())
	}

	union = numbers_array.Element.(domain.UnionType)

	if union.OneOf[0].Name() != (domain.IntType{}).Name() {
		test.Fatalf("Wrong union[0] type expected %s, received %s", (domain.IntType{}).Name(), numbers_array.Element.Name())
	}
	if union.OneOf[1].Name() != (domain.FloatType{}).Name() {
		test.Fatalf("Wrong union[1] type expected %s, received %s", (domain.FloatType{}).Name(), numbers_array.Element.Name())
	}

	empty_array := graph.Fields["empty"].(domain.ArrayType)

	if empty_array.Element.Name() != (domain.UnknownType{}).Name() {
		test.Fatalf("Wrong array element type expected %s, received %s", (domain.UnknownType{}).Name(), empty_array.Element.Name())
	}
}
