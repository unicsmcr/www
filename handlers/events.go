package handlers

import (
	"net/http"
	"github.com/alexdmtr/www/services/eventService"
)

func events(w http.ResponseWriter, r *http.Request) {

	events := eventService.GetEvents()

	templates["events"].ExecuteTemplate(w, "layout", &events)
}