package jsdoc

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
	"mixed": [1, "2", false, {"id": 3, "name": "Simon"}],
	"float": 69.420,
	"cool": true
	}
	`

	lexer := json.New(input)
	parser := parser.New(lexer)

	graph := graph.New(dedup.New())
	codegen := New()
	result := codegen.Generate(graph.Generate(parser.Parse()))
	expected := `/**
 * @typedef  { Object } User
 * @property { number } id
 * @property { string } name
 */

/**
 * @typedef  { Object } Root
 * @property { number } id
 * @property { User } author
 * @property { boolean } cool
 * @property { number } float
 * @property { (User | boolean | number | string)[] } mixed
 * @property { number[] } numbers
 * @property { string } title
 * @property { User } user
 */
`

	errors := utils.CompareLineByLine(result, expected)
	if len(errors) != 0 {
		test.Fatal(strings.Join(errors, "\n\n"))
	}

}
