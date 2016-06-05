package handlers

import (
	"net/http"
	"www/services/eventService"	
)

type indexModel struct {
    HasUpcomingEvent bool
    Event *eventService.Event
}

func index(w http.ResponseWriter, r *http.Request) {
	var model indexModel
	
	model.Event = eventService.GetUpcomingEvent()
	model.HasUpcomingEvent = model.Event != nil
	templates["index"].ExecuteTemplate(w, "layout", &model)
}