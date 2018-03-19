package handlers

import (
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	var message messageModel
	switch status {
	case http.StatusNotFound:
		message = messageModel{"Error", "Page not found."}
	case http.StatusServiceUnavailable:
		message = messageModel{"Error", "Page not available at this time. Try again later."}
	default:
		message = messageModel{"Error", "An error has occurred."}
	}

	err := templates["message"].ExecuteTemplate(w, "layout", message)
	if err != nil {
		log.Println(err)
	}
}
