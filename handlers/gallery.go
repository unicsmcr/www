package handlers

import "net/http"

func gallery(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/gallery" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}