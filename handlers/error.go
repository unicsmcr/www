package handlers

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	
	if status == http.StatusNotFound {
		templates["error"].ExecuteTemplate(w, "layout", "Page not found.")
	} else {
		templates["error"].ExecuteTemplate(w, "layout", "An error has occurred.")		
	}
}