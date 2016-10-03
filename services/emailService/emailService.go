package emailService

import (
	"errors"
	"github.com/hacksoc-manchester/www/helpers/validator"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

func trySendEmail(senderName, senderEmail, receiverName, receiverEmail, subject, message string) bool {
	from := mail.NewEmail(senderName, senderEmail)
	to := mail.NewEmail(receiverName, receiverEmail)
	content := mail.NewContent("text/plain", message)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)

	response, err := sendgrid.API(request)

	return err != nil || response.StatusCode == 202
}

// Send transmits an email to the specified receiver.
func Send(senderName, senderEmail, receiverName, receiverEmail, subject, message string) error {
	if !validator.IsValidName(senderName) {
		return errors.New(`Name "` + senderName + `" is not valid.`)
	}

	if !validator.IsValidEmail(senderEmail) {
		return errors.New(`Email address "` + senderEmail + `" is not valid.`)
	}

	if !validator.IsValidName(receiverName) {
		return errors.New(`Name "` + receiverName + `" is not valid.`)
	}

	if !validator.IsValidEmail(receiverEmail) {
		return errors.New(`Email address "` + receiverEmail + `" is not valid.`)
	}

	if !validator.IsValidMessage(message) {
		return errors.New("Please provide a message.")
	}

	if !trySendEmail(senderName, senderEmail, receiverName, receiverEmail, subject, message) {
		return errors.New("An unexpected error has occurred.")
	}

	return nil
}
