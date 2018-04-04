package handlers

import (
	"log"
	"net/http"

	"github.com/hacksoc-manchester/www/services/eventService"
)

func events(w http.ResponseWriter, r *http.Request) {

	var eventsContext struct {
		EventGroup   *eventService.EventGroup
		HaveOngoing  bool
		HaveUpcoming bool
	}
	eventGroup, err := eventService.GroupEvents()

	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusServiceUnavailable)
		return
	}

	eventsContext.EventGroup = eventGroup
	eventsContext.HaveOngoing = len(eventGroup.Ongoing) > 0
	eventsContext.HaveUpcoming = len(eventGroup.Upcoming) > 0

	renderTemplate(w, r, "events", eventsContext)
}
