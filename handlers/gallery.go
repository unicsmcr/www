package handlers

import "net/http"

func gallery(w http.ResponseWriter, r *http.Request) {
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}