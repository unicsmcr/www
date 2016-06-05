package handlers

import "net/http"

func contact(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}