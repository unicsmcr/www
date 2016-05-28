package handlers

import "net/http"

type indexModel struct {
    HasUpcomingEvent bool
    EventName string
    EventDate string
    Menu []struct {
        Name string
        Tag string
    }
}

func index(w http.ResponseWriter, r *http.Request) {
	var model indexModel
	
	model.HasUpcomingEvent = true
	model.EventName = "GreatUniHack 2016"
	model.EventDate = "September 16 at 5:30pm"
	model.Menu = []struct {
		Name string
		Tag string
	}{
		{"Events", "events"},
		{"Team", "team"},
		{"Gallery", "gallery"},
		{"Contact", "contact"},
	}
  	
	templates["index"].ExecuteTemplate(w, "layout", &model)
}