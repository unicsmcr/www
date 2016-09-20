package handlers

import (
	"net/http"
)

func newsletter(w http.ResponseWriter, r *http.Request) {
	// TODO(andrei): Finish the newsletter page.
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}