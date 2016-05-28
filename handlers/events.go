package handlers

import "net/http"

func events(w http.ResponseWriter, r *http.Request) {
	templates["events"].ExecuteTemplate(w, "layout", nil)
}