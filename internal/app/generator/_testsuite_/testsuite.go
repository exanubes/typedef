package testsuite

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/generator"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
	"github.com/exanubes/typedef/internal/utils"
)

func CodegenTestSuite(test *testing.T, name string, codegen generator.CodeGenerator) {
	test.Run(name, func(test *testing.T) {
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
		result := strings.TrimSpace(codegen.Generate(graph.Generate(parser.Parse())))
		expected := strings.TrimSpace(read_stub_file(name))
		errors := utils.CompareLineByLine(expected, result)
		if len(errors) != 0 {
			test.Fatal(strings.Join(errors, "\n\n"))
		}
	})
}

func read_stub_file(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "failed to get caller information"
	}

	dir := filepath.Dir(filename)
	testdataPath := filepath.Join(dir, "testdata", name, "named_types.txt")

	expected, err := os.ReadFile(testdataPath)
	if err != nil {
		return err.Error()
	}

	return string(expected)
}
