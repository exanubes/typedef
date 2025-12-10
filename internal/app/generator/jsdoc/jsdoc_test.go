package jsdoc_test

import (
	"testing"

	testsuite "github.com/exanubes/typedef/internal/app/generator/_testsuite_"
	"github.com/exanubes/typedef/internal/app/generator/jsdoc"
)

func TestJsdocCodegenSuite(test *testing.T) {
	testsuite.CodegenTestSuite(test, "jsdoc", jsdoc.New())
}
