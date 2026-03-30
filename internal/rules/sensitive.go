package rules

import (
	"strings"
)

var sensitiveKeywords = []string{
	"password",
	"token",
	"api_key",
	"apikey",
	"secret",
	"credentials",
}

func CheckSensitive(message string) string {
	lowerMsg := strings.ToLower(message)

	for _, word := range sensitiveKeywords {
		if strings.Contains(lowerMsg, word) {
			return "log messages should not contain potentially sensitive data"
		}
	}

	return ""
}
