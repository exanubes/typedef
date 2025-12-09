package generator_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/exanubes/typedef/internal/app/dedup"
	"github.com/exanubes/typedef/internal/app/generator"
	"github.com/exanubes/typedef/internal/app/generator/golang"
	"github.com/exanubes/typedef/internal/app/generator/jsdoc"
	"github.com/exanubes/typedef/internal/app/generator/typescript"
	"github.com/exanubes/typedef/internal/app/generator/zod"
	"github.com/exanubes/typedef/internal/app/graph"
	"github.com/exanubes/typedef/internal/app/lexer/json"
	parser "github.com/exanubes/typedef/internal/app/parser/json"
	"github.com/exanubes/typedef/internal/utils"
)

type CodegenUnderTest struct {
	Name   string
	Create func() generator.CodeGenerator
}

func TestCodegenSuite(test *testing.T) {
	factories := []CodegenUnderTest{
		{
			Name:   "golang",
			Create: func() generator.CodeGenerator { return golang.New(func(_ int) string { return "testing" }) },
		},

		{
			Name:   "typescript",
			Create: func() generator.CodeGenerator { return typescript.New() },
		},
		{
			Name:   "zod",
			Create: func() generator.CodeGenerator { return zod.New() },
		},
		{
			Name:   "jsdoc",
			Create: func() generator.CodeGenerator { return jsdoc.New() },
		},
	}

	for _, factory := range factories {
		test.Run(factory.Name, func(test *testing.T) {
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
			codegen := factory.Create()
			result := strings.TrimSpace(codegen.Generate(graph.Generate(parser.Parse())))
			expected := strings.TrimSpace(read_stub_file(factory.Name))
			errors := utils.CompareLineByLine(expected, result)
			if len(errors) != 0 {
				test.Fatal(strings.Join(errors, "\n\n"))
			}
		})
	}
}

func read_stub_file(name string) string {
	expected, err := os.ReadFile(fmt.Sprintf("_testdata_/%s/named_types.txt", name))

	if err != nil {
		return err.Error()
	}

	return string(expected)
}
