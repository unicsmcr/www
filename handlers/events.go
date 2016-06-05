package handlers

import (
	"net/http"
	"www/services/eventService"
)

func events(w http.ResponseWriter, r *http.Request) {
	events := eventService.GetEvents()
	
	templates["events"].ExecuteTemplate(w, "layout", &events)
}