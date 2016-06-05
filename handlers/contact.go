package handlers

import (
	"os"
	"net/http"
)

var reCaptchaSiteKey = os.Getenv("RECAPTCHA_SITE_KEY")

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates["contact"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)
			
		case "POST":
			templates["contact"].ExecuteTemplate(w, "layout", reCaptchaSiteKey)			
			
		default:
			errorHandler(w, r, http.StatusBadRequest)
	}
}