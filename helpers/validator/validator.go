package validator

import (
	"regexp"
	"strings"
)

// IsValidEmail determines whether the specified email is valid.
func IsValidEmail(email string) bool {
	result, _ := regexp.MatchString(`^[^ @]+@[^ @]+\.[^ @]+$`, email)

	return len(email) <= 100 && result
}

// IsValidMessage determines whether the specified message is valid.
func IsValidMessage(message string) bool {
	return len(message) <= 4000 && len(strings.Replace(message, " ", "", -1)) > 0
}

// IsValidName determines whether the specified name is valid.
func IsValidName(name string) bool {
	return len(name) <= 30 && len(strings.Replace(name, " ", "", -1)) > 0
}
