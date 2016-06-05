package handlers

import "net/http"

func contact(w http.ResponseWriter, r *http.Request) {
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}