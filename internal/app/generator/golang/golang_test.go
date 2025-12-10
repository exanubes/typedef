package golang_test

import (
	"testing"

	testsuite "github.com/exanubes/typedef/internal/app/generator/_testsuite_"
	"github.com/exanubes/typedef/internal/app/generator/golang"
)

func TestGolangCodegenSuite(test *testing.T) {
	testsuite.CodegenTestSuite(test, "golang", golang.New(func(_ int) string { return "testing" }))
}
