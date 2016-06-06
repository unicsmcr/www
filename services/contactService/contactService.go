package contactService

import (
	"errors"
	"os"
	"regexp"
	"github.com/sendgrid/sendgrid-go"
	"strings"
)

var contactEmail = os.Getenv("CONTACT_EMAIL")
var username = os.Getenv("SENDGRID_USERNAME")
var password = os.Getenv("SENDGRID_PASSWORD")
var emailer = sendgrid.NewSendGridClient(username, password)

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

func sendMessage(senderName, senderEmail, message string) error {
	mail := sendgrid.NewMail()
	
	mail.AddTo(contactEmail)
	mail.SetFrom(senderEmail)
	mail.SetSubject(senderName)
	mail.SetText(message)
	
	return emailer.Send(mail)
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
	
	if err := sendMessage(senderName, senderEmail, message); err != nil {
		return errors.New("An unexpected error has occurred.")
	}
	
	return nil
}