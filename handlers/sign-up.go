package handlers

import (
	"errors"
	"github.com/hacksoc-manchester/www/helpers/crypto"
	"github.com/hacksoc-manchester/www/services/databaseService"
	"github.com/hacksoc-manchester/www/services/emailService"
	"net/http"
	"os"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates["sign-up"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)

	case "POST":
		firstName := r.PostFormValue("first-name")
		lastName := r.PostFormValue("last-name")
		email := r.PostFormValue("email")
		subscribedToArticles := r.PostFormValue("subscribe-to-articles") == "on"
		subscribedToEvents := r.PostFormValue("subscribe-to-events") == "on"

		var response string

		if reCaptcha.Verify(*r) {
			err := registerUser(firstName, lastName, email, subscribedToArticles, subscribedToEvents)

			if err == nil {
				response = "Welcome! You are now part of our mailing list."
			} else {
				response = err.Error()
			}
		} else {
			response = "Turing test failed. Please try again."
		}

		templates["message"].ExecuteTemplate(w, "layout", messageModel{"Sign Up", response})

	default:
		errorHandler(w, r, http.StatusBadRequest)
	}
}

func registerUser(firstName, lastName, email string, subscribedToArticles, subscribedToEvents bool) error {
	if !subscribedToArticles && !subscribedToEvents {
		return errors.New("Please select at least one subscription.")
	}

	err := databaseService.CreateUser(firstName, lastName, email, subscribedToArticles, subscribedToEvents)

	if err != nil {
		return err
	}

	senderName := "HackSoc"
	senderEmail := os.Getenv("NOREPLY_EMAIL")
	receiverName := firstName + " " + lastName
	subject := "Welcome to HackSoc!"
	message := `You are now part of our mailing list.
		<br>
		<br>
		To unsubscribe, <a href=\"` + getUnsubscribeLink(email) + `">click here</a>.`

	return emailService.Send(senderName, senderEmail, receiverName, email, subject, message)
}

func getUnsubscribeLink(email string) string {
	token, _ := crypto.Encrypt(email)

	return "http://hacksoc.com/unsubscribe?token=" + token
}
