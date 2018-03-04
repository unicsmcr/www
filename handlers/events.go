package handlers

import (
	"github.com/alexdmtr/www/services/eventService"
	"log"
	"net/http"
)

func events(w http.ResponseWriter, r *http.Request) {

	var eventsContext struct {
		EventGroup   *eventService.EventGroup
		HaveRightNow bool
		HaveUpcoming bool
	}
	eventGroup, err := eventService.GroupEvents()

	if err != nil {
		log.Println(err)
	}

	eventsContext.EventGroup = eventGroup
	eventsContext.HaveRightNow = len(eventGroup.RightNow) > 0
	eventsContext.HaveUpcoming = len(eventGroup.Upcoming) > 0

	err = templates["events"].ExecuteTemplate(w, "layout", eventsContext)

	if err != nil {
		log.Println(err)
	}
}
