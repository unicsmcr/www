package handlers

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	
	if status == http.StatusNotFound {
		templates["message"].ExecuteTemplate(w, "layout", messageModel{"Error", "Page not found."})
	} else {
		templates["message"].ExecuteTemplate(w, "layout", messageModel{"Error", "An error has occurred."})
	}
}