package handlers

import (
	"net/http"
	"os"
	"github.com/hacksoc-manchester/www/services/contactService"
	"github.com/haisum/recaptcha"
)

var reCaptchaSiteKey = os.Getenv("RECAPTCHA_SITE_KEY")
var re = recaptcha.R { 
	Secret: os.Getenv("RECAPTCHA_SECRET_KEY"),
}

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates["contact"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)
			
		case "POST":
			senderName := r.PostFormValue("sender-name")
			senderEmail := r.PostFormValue("sender-email")
			message := r.PostFormValue("message")
			var response string
			
			if re.Verify(*r) {
				response = contactService.Send(senderName, senderEmail, message)				
			} else {
				response = "Turing test failed. Please try again!"
			}
			
			templates["message"].ExecuteTemplate(w, "layout", messageModel{"Contact", response})
			
		default:
			errorHandler(w, r, http.StatusBadRequest)
	}
}