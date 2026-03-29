package rules

import (
	"unicode"
	"unicode/utf8"
)

func CheckLowercase(message string) string {
	if message == "" {
		return ""
	}

	r, _ := utf8.DecodeRuneInString(message)
	if unicode.IsUpper(r) {
		return "log messages must begin with a lowercase letter"
	}

	return ""
}
