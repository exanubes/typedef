package golang

import (
	"strings"
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
	"github.com/exanubes/typedef/internal/utils"
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
	"float": 69.420,
	"cool": true
	}
	`

	lexer := json.New(input)
	parser := parser.New(lexer)

	graph := graph.New(dedup.New())
	codegen := New()
	result := codegen.Generate(graph.Generate(parser.Parse()))
	// TODO: mock utils.RandomString() to be able to test union types
	expected := `type User struct {
  ID int
  Name string
}

type Root struct {
  ID int
  Author User
  Cool bool
  Float float64
  Numbers []int
  Title string
  User User
}
`

	errors := utils.CompareLineByLine(result, expected)
	if len(errors) != 0 {
		test.Fatal(strings.Join(errors, "\n\n"))
	}

}
