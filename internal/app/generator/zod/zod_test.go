package zod_test

import (
	"testing"

	testsuite "github.com/exanubes/typedef/internal/app/generator/_testsuite_"
	"github.com/exanubes/typedef/internal/app/generator/zod"
)

func TestZodCodegenSuite(test *testing.T) {
	testsuite.CodegenTestSuite(test, "zod", zod.New())
}
