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
			if !re.Verify(*r) {
				const response = "Turing test failed. Please try again!"

				templates["message"].ExecuteTemplate(w, "layout", messageModel{"Contact", response})
				
				return
			}
		
			senderName := r.FormValue("sender-name")
			senderEmail := r.FormValue("sender-email")
			message := r.FormValue("message")
			response := contactService.Send(senderName, senderEmail, message)
			
			templates["message"].ExecuteTemplate(w, "layout", messageModel{"Contact", response})
			
		default:
			errorHandler(w, r, http.StatusBadRequest)
	}
}