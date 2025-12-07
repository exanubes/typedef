package golang

import "strings"

func capitalize(value string) string {
	if strings.ToLower(value) == "id" {
		return "ID"
	}

	return strings.ToUpper(value[:1]) + value[1:]
}
