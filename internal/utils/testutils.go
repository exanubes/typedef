package utils

import (
	"fmt"
	"strings"
)

const (
	red   = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

func CompareLineByLine(result, expected string) []string {
	result_lines := strings.Split(result, "\n")
	expected_lines := strings.Split(expected, "\n")
	errors := []string{}
	for index, line := range result_lines {
		if line != expected_lines[index] {
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("Mismatch on line %d", index))
			builder.WriteRune('\n')
			builder.WriteString(green + "Expected:" + reset)
			builder.WriteRune('\n')
			builder.WriteString(fmt.Sprintf("%d |  %s", index, expected_lines[index]))
			builder.WriteRune('\n')
			builder.WriteString(red + "Received:" + reset)
			builder.WriteRune('\n')
			builder.WriteString(fmt.Sprintf("%d |  %s", index, line))
			builder.WriteRune('\n')
			errors = append(errors, builder.String())
		}
	}

	return errors
}
