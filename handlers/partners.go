package handlers

import "net/http"

func partners(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/partners" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}