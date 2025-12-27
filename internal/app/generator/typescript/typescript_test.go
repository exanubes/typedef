//go:build !integration

package typescript_test

import (
	"testing"

	testsuite "github.com/exanubes/typedef/internal/app/generator/_testsuite_"
	"github.com/exanubes/typedef/internal/app/generator/typescript"
)

func TestTypescriptCodegenSuite(test *testing.T) {
	testsuite.CodegenTestSuite(test, "typescript", typescript.New())
}
