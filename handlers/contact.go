package handlers

import (
	"net/http"
	"os"

	"github.com/hacksoc-manchester/www/helpers/validator"
	"github.com/hacksoc-manchester/www/services/emailService"
)

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates["contact"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)

	case "POST":
		senderName := r.PostFormValue("name")
		senderEmail := r.PostFormValue("email")
		message := r.PostFormValue("message")
		var response string

		if reCaptcha.Verify(*r) {
			err := contactHackSoc(senderName, senderEmail, message)

			if err == nil {
				response = "Your message has been received."
			} else {
				response = err.Error()
			}
		} else {
			response = "Turing test failed. Please try again."
		}

		renderTemplate(w, r, "message", messageModel{"Contact", response})

	default:
		errorHandler(w, r, http.StatusBadRequest)
	}
}

func contactHackSoc(senderName, senderEmail, message string) error {
	if !validator.IsValidName(senderName) {
		senderName = "Anonymous"
	}

	if senderEmail == "" {
		senderEmail = os.Getenv("NOREPLY_EMAIL")
	}

	receiverName := "HackSoc"
	receiverEmail := os.Getenv("CONTACT_EMAIL")
	subject := "Contact Form Message"

	return emailService.Send(senderName, senderEmail, receiverName, receiverEmail, subject, message)
}
