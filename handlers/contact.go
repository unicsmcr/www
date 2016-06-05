package handlers

import (
	"net/http"
	"os"
	"github.com/hacksoc-manchester/www/services/contactService"
)

var reCaptchaSiteKey = os.Getenv("RECAPTCHA_SITE_KEY")

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates["contact"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)
			
		case "POST":
			senderName := r.FormValue("sender-name")
			senderEmail := r.FormValue("sender-email")
			message := r.FormValue("message")
			response := contactService.Send(senderName, senderEmail, message)
			
			templates["messagesent"].ExecuteTemplate(w, "layout", response)
			
		default:
			errorHandler(w, r, http.StatusBadRequest)
	}
}