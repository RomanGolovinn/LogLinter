package rules

import (
	"strings"
	"unicode"
)

func CheckSymbols(message string) string {
	if strings.ContainsAny(message, "!?") || strings.Contains(message, "...") {
		return "log messages must not contain special characters."
	}

	for _, r := range message {
		if unicode.IsSymbol(r) || unicode.IsMark(r) {
			return "log messages must not contain special characters."
		}
	}

	return ""
}
