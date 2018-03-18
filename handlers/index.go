package handlers

import (
	"net/http"

	"github.com/alexdmtr/www/services/eventService"
)

type indexModel struct {
	HasUpcomingEvent bool
	Event            *eventService.Event
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	var model indexModel

	model.HasUpcomingEvent = model.Event != nil
	renderTemplate(w, r, "index", &model)
}
