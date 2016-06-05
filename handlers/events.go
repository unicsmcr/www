package handlers

import (
	"net/http"
	"github.com/hacksoc-manchester/www/services/eventService"
)

func events(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/events" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	
	events := eventService.GetEvents()
	
	templates["events"].ExecuteTemplate(w, "layout", &events)
}