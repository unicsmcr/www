package contactService

import (
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	"regexp"
	"strings"
)

func emailIsValid(email string) bool {
	result, _ := regexp.MatchString(`^[^ @]+@[^ @]+\.[^ @]+$`, email)
	
	return result
}

func messageIsValid(message string) bool {
	return len(strings.Replace(message, " ", "", -1)) > 0
}

func nameIsValid(name string) bool {
	return len(strings.Replace(name, " ", "", -1)) > 0	
}

func trySendEmail(senderName, senderEmail, message string) bool {
	from := mail.NewEmail(senderName, senderEmail)
	subject := "Contact Form Message"
	to := mail.NewEmail("HackSoc", os.Getenv("CONTACT_EMAIL"))
	content := mail.NewContent("text/plain", message)
	m := mail.NewV3MailInit(from, subject, to, content)
	
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)

	response, err := sendgrid.API(request)

	return err != nil || response.StatusCode == 202
}

// Send sends an email to HackSoc.
func Send(senderName, senderEmail, message string) error {
	if !nameIsValid(senderName) {
		senderName = "Anonymous"
	}
	
	if senderEmail == "" {
		senderEmail = os.Getenv("NOREPLY_EMAIL")
	} else if !emailIsValid(senderEmail) {
		return errors.New("Email address \"" + senderEmail + "\" is not valid.")
	}
	
	if !messageIsValid(message) {
		return errors.New("Please provide a message.")
	}

	if !trySendEmail(senderName, senderEmail, message) {
		return errors.New("An unexpected error has occurred.")
	}
	
	return nil
}