package graph

import "time"

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
