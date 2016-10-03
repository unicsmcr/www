package newsletterService

import (
	"errors"
	"regexp"
)

func emailIsValid(email string) bool {
	result, _ := regexp.MatchString(`^[^ @]+@[^ @]+\.[^ @]+$`, email)

	return result
}

// SubscribeToArticles subscribes the user to articles.
func SubscribeToArticles(email string) error {
	if !emailIsValid(email) {
		return errors.New("Email address \"" + email + "\" is not valid.")
	}

	return nil
}

// SubscribeToEvents subscribes the user to events.
func SubscribeToEvents(email string) error {
	if !emailIsValid(email) {
		return errors.New("Email address \"" + email + "\" is not valid.")
	}

	return nil
}
