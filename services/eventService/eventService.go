package eventService

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// An Event represents a HackSoc event.
type Event struct {
	Name           string
	Description    string
	Location       string
	URL            string
	ImageURL       string
	Upcoming       bool
	AttendingCount int
	StartTime      time.Time
	EndTime        time.Time
}

const fbGraphApiUrl = "https://graph.facebook.com/oauth/access_token"

var accessToken string

var eventsCache *cache.Cache

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
	eventsCache = cache.New(10*time.Minute, 10*time.Minute)

	eventsCache.Set("events", getEvents(), cache.DefaultExpiration)
	log.Println("Save in cache.")
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			eventsCache.Set("events", getEvents(), cache.DefaultExpiration)
			log.Println("Update cache!")
		}
	}()
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

// getEvents gets the events.
func getEvents() []*Event {

	req := requestWithData("GET",
		"https://graph.facebook.com/v2.12/hacksocmcr/events", map[string]string{
			"fields": "name,description,cover,place,start_time,end_time,attending_count",
		}, true)

	client := &http.Client{}
	response, _ := client.Do(req)

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	type EventEntry struct {
		Name        string `json:"name"`
		Description string `json:"description"`

		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`

		AttendingCount int `json:"attending_count"`

		Cover struct {
			Source string `json:"source"`
		} `json:"cover"`
		Place struct {
			Name string `json:"name"`
		} `json:"place"`
	}

	type BodyData struct {
		Data []EventEntry `json:"data"`
	}

	var data BodyData
	err := json.Unmarshal(body, &data)

	if err != nil {
		log.Println(err)
	}

	events := make([]*Event, 0)

	for _, event := range data.Data {
		layout := "2006-01-02T15:04:05-0700"
		startTime, err := time.Parse(layout, event.StartTime)
		if err != nil {
			log.Println(err)
		}

		endTime, err := time.Parse(layout, event.EndTime)
		if err != nil {
			log.Println(err)
		}

		today := time.Now()
		upcoming := startTime.After(today)

		events = append(events, &Event{
			event.Name,
			event.Description,
			event.Place.Name,
			"",
			event.Cover.Source,
			upcoming,
			event.AttendingCount,
			startTime,
			endTime,
		})
	}

	return events
}

type EventGroup struct {
	RightNow []*Event
	Upcoming []*Event
	Past     []*Event
}

func getEventsFromCache() ([]*Event, bool) {
	events, found := eventsCache.Get("events")
	log.Println("Get from cache!")

	return events.([]*Event), found
}

// GroupEvents groups the events by 'right now', 'upcoming', and 'past'
func GroupEvents() (*EventGroup, error) {
	events, _ := getEventsFromCache()

	eventGroup := &EventGroup{}
	eventGroup.RightNow = make([]*Event, 0)
	eventGroup.Upcoming = make([]*Event, 0)
	eventGroup.Past = make([]*Event, 0)

	for _, event := range events {
		// (todo:Alex) Add "RightNow" events
		if event.Upcoming {
			eventGroup.Upcoming = append(eventGroup.Upcoming, event)
		} else {
			eventGroup.Past = append(eventGroup.Past, event)
		}
	}

	return eventGroup, nil
}
