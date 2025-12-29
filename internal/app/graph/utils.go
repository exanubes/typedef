package graph

import (
	"regexp"
	"time"
)

func is_date_string(input string) bool {
	if input == "" {
		return false
	}

	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		time.DateOnly,
		time.DateTime,
	}

	for _, format := range formats {
		if _, err := time.Parse(format, input); err == nil {
			return true
		}
	}

	return false
}

var uuid_regexe = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func is_uuid_string(input string) bool {
	return uuid_regexe.MatchString(input)
}
