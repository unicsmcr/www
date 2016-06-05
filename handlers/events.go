package handlers

import (
	"net/http"
	"github.com/hacksoc-manchester/www/services/eventService"
)

func events(w http.ResponseWriter, r *http.Request) {
	events := eventService.GetEvents()
	
	templates["events"].ExecuteTemplate(w, "layout", &events)
}