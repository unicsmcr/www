package handlers

import "net/http"

func newsletter(w http.ResponseWriter, r *http.Request) {
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}