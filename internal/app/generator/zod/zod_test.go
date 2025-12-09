package zod

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
	expected := `import { z } from "zod";

const UserSchema = z.object({
  id: z.number(),
  name: z.string(),
});
type User = z.infer<typeof UserSchema>;

const RootSchema = z.object({
  id: z.number(),
  author: UserSchema,
  cool: z.boolean(),
  float: z.number(),
  mixed: z.array(z.union([UserSchema, z.boolean(), z.number(), z.string()])),
  numbers: z.array(z.number()),
  title: z.string(),
  user: UserSchema,
});
type Root = z.infer<typeof RootSchema>;
`
	errors := utils.CompareLineByLine(result, expected)
	if len(errors) != 0 {
		test.Fatal(strings.Join(errors, "\n\n"))
	}

}
