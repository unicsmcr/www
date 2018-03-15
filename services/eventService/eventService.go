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
	Name            string
	Description     string
	Location        string
	URL             string
	ImageURL        string
	AttendingCount  int
	InterestedCount int
	StartTime       time.Time
	EndTime         time.Time
}

func (e *Event) GetInterestLine() string {
	if e.StartTime.After(time.Now()) || e.EndTime.After(time.Now()) {
		totalCount := e.AttendingCount + e.InterestedCount
		return fmt.Sprintf("%d going or interested", totalCount)
	}
	return fmt.Sprintf("%d attended", e.AttendingCount)
}

func (e *Event) GetShortDate() string {
	if e.StartTime.Year() == time.Now().Year() && e.StartTime.After(time.Now()) {
		startTime := e.StartTime.Format("January 2")
		if e.StartTime.YearDay() != e.EndTime.YearDay() {
			if e.StartTime.Month() == e.EndTime.Month() {
				return fmt.Sprintf("%s - %s", startTime, e.EndTime.Format("2"))
			}
			var endTime string

			if e.EndTime.Year() == time.Now().Year() {
				endTime = e.EndTime.Format("January 2")
			} else {
				endTime = e.EndTime.Format("January 2, 2006")
			}

			return fmt.Sprintf("%s - %s", startTime, endTime)
		} else {
			return startTime
		}
	}
	return e.StartTime.Format("January 2, 2006")
}

const timeFormatLayout = "2006-01-02T15:04:05-0700"

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

const MAX_SHORTNAME_LENGTH = 30

// getEvents gets the events.
func getEvents() []*Event {

	req := requestWithData("GET",
		"https://graph.facebook.com/v2.12/hacksocmcr/events", map[string]string{
			"fields": "name,description,cover,place,start_time,end_time,attending_count,interested_count",
		}, true)

	client := &http.Client{}
	response, _ := client.Do(req)

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	type EventEntry struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`

		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`

		AttendingCount  int `json:"attending_count"`
		InterestedCount int `json:"interested_count"`

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
		startTime, err := time.Parse(timeFormatLayout, event.StartTime)

		if err != nil {
			log.Println(err)
		}

		endTime, err := time.Parse(timeFormatLayout, event.EndTime)
		if err != nil {
			log.Println(err)
		}

		events = append(events, &Event{
			event.Name,
			event.Description,
			event.Place.Name,
			"https://facebook.com/" + event.ID,
			event.Cover.Source,
			event.AttendingCount,
			event.InterestedCount,
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

	now := time.Now()

	for _, event := range events {
		upcoming := event.StartTime.After(now)
		rightNow := event.StartTime.Before(now) && event.EndTime.After(now)

		if rightNow {
			eventGroup.RightNow = append(eventGroup.RightNow, event)
		} else if upcoming {
			eventGroup.Upcoming = append(eventGroup.Upcoming, event)
		} else {
			eventGroup.Past = append(eventGroup.Past, event)
		}
	}

	return eventGroup, nil
}
