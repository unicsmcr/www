package handlers

import "net/http"

type indexModel struct {
    HasUpcomingEvent bool
    EventName string
    EventDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	var model indexModel
	
	model.HasUpcomingEvent = true
	model.EventName = "GreatUniHack 2016"
	model.EventDate = "September 16 at 5:30pm"
	templates["index"].ExecuteTemplate(w, "layout", &model)
}