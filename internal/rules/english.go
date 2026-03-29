package rules

import (
	"unicode"
)

func CheckEnglish(message string) string {
	for _, r := range message {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return "log messages must be in english only."
		}
	}

	return ""
}
