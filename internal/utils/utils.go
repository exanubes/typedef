package utils

import (
	"crypto/rand"
	"slices"
	"strings"
)

func Capitalize(value string) string {
	if strings.ToLower(value) == "id" {
		return "ID"
	}

	return strings.ToUpper(value[:1]) + value[1:]
}

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345667890-"
const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Letter(index int) string {
	return string(alpha[index%len(alpha)])
}

func RandomString(length int) string {

	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = alphanumeric[int(b[i])%len(alphanumeric)]
	}
	return string(b)
}

func SortFields(fields []string) []string {
	slices.SortStableFunc(fields, func(a, b string) int {
		if strings.ToLower(a) == "id" {
			return -1
		}

		if strings.ToLower(b) == "id" {
			return 1
		}

		if strings.ToLower(a) > strings.ToLower(b) {
			return 1
		}

		if strings.ToLower(a) < strings.ToLower(b) {
			return -1
		}

		return 0
	})

	return fields
}
