package eventService

// An Event represents a HackSoc event.
type Event struct {
	Name        string
	Description string
	Location    string
	Date        string
	URL         string
	ImageURL    string
	Upcoming    bool
}

// GetEvents gets the events.
func GetEvents() []*Event {
	events := []*Event{
		&Event{"GreatUniHack 2016", "A first test", "Kilburn Building", "September 16 at 5:30pm", "https://google.com/", "https://www.google.co.uk/images/branding/product/ico/googleg_lodp.ico", true},
		&Event{"Test 2", "A second test", "Here", "June 1 at 12:00pm", "https://google.com/", "https://www.google.co.uk/images/branding/product/ico/googleg_lodp.ico", false},
		&Event{"Test 3", "A third test", "Here", "May 25 at 10:00am", "https://google.com/", "https://www.google.co.uk/images/branding/product/ico/googleg_lodp.ico", false},
	}

	return events
}

// GetUpcomingEvents gets the upcoming events.
func GetUpcomingEvents() []*Event {
	var upcomingEvents []*Event

	for _, event := range GetEvents() {
		if event.Upcoming {
			upcomingEvents = append(upcomingEvents, event)
		}
	}

	return upcomingEvents
}

// GetUpcomingEvent gets the earliest upcoming event.
func GetUpcomingEvent() *Event {
	upcomingEvents := GetUpcomingEvents()

	return upcomingEvents[len(upcomingEvents)-1]
}
