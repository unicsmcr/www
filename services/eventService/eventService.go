package eventService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

const fbGraphApiUrl = "https://graph.facebook.com/oauth/access_token"

var accessToken string

func init() {
	if os.Getenv("FB_APP_ID") == "" {
		log.Println("Environment variable FB_APP_ID is not assigned.")
		return
	}

	if os.Getenv("FB_SECRET") == "" {
		log.Println("Environment variable FB_SECRET is not assigned.")
		return
	}

	accessToken = getAccessToken()
}

// getAccessToken gets a suitable Facebook GraphAPI access token.
func getAccessToken() string {
	accessToken := fmt.Sprintf("%s|%s", os.Getenv("FB_APP_ID"), os.Getenv("FB_SECRET"))

	return accessToken
}

func requestWithData(method, url string, data map[string]string, useToken bool) *http.Request {
	request, _ := http.NewRequest(method, url, nil)

	query := request.URL.Query()
	for key, value := range data {
		query.Add(key, value)
	}

	if useToken {
		request.Header.Add("Authorization", "Bearer "+accessToken)
	}

	request.URL.RawQuery = query.Encode()

	return request
}

// GetEvents gets the events.
func GetEvents() []*Event {

	req := requestWithData("GET",
		"https://graph.facebook.com/v2.12/hacksocmcr/events", map[string]string{
			"fields": "name,description,cover,place,start_time",
		}, true)

	client := &http.Client{}
	response, _ := client.Do(req)

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	type EventEntry struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Cover       struct {
			Source string `json:"source"`
		} `json:"cover"`
		Place struct {
			Name string `json:"name"`
		} `json:"place"`
		StartTime string `json:"start_time"`
	}

	type BodyData struct {
		Data []EventEntry `json:"data"`
	}

	var data BodyData
	err := json.Unmarshal(body, &data)

	if err != nil {
		log.Println(err)
	}

	events := []*Event{}

	for _, event := range data.Data {
		events = append(events, &Event{
			event.Name,
			event.Description,
			event.Place.Name,
			event.StartTime,
			"",
			event.Cover.Source,
			false})
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
