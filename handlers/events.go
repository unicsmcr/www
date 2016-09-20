package handlers

import (
	"net/http"
)

func events(w http.ResponseWriter, r *http.Request) {
	// TODO(andrei): Finish the events page.
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}