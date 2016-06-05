package handlers

import "net/http"

func newsletter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newsletter" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}