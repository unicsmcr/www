package handlers

import (
	"net/http"
	"github.com/hacksoc-manchester/www/services/newsletterService"
)

func newsletter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates["newsletter"].ExecuteTemplate(w, "layout", nil)
			
		case "POST":
			events := r.PostFormValue("events") == "on"
			articles := r.PostFormValue("articles") == "on"
			email := r.PostFormValue("email")
			var err error
			
			// Subscribes the user.
			if events {
				err = newsletterService.SubscribeToEvents(email)
			}
			
			if err == nil && articles {
				err = newsletterService.SubscribeToArticles(email)
			}
			
			response := "This feature is currently unavailable."

			if err != nil {
				response = err.Error()
			}
			
			templates["message"].ExecuteTemplate(w, "layout", messageModel{"Newsletter", response})			
			
		default:
			errorHandler(w, r, http.StatusBadRequest)
	}
}