package parser

import (
	"testing"

	"github.com/exanubes/typedef/internal/app/lexer/json"
)

func TestNextToken(test *testing.T) {
	input := `
	{
	"property": true
	}`

	lexer := json.New(input)
	parser := New(lexer)

	program := parser.Parse()

	print("object", program.Value.String())

	if true != false {
	} else {

		test.Fatal("Failed")
	}

}
