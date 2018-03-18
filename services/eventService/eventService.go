package eventService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexdmtr/www/config"
	"github.com/patrickmn/go-cache"
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
			endTime := e.EndTime.Format("January 2, 2006")
			if e.EndTime.Year() == time.Now().Year() {
				endTime = e.EndTime.Format("January 2")
			}
			return fmt.Sprintf("%s - %s", startTime, endTime)
		}
		return startTime
	}
	return e.StartTime.Format("January 2, 2006")
}

const timeFormatLayout = "2006-01-02T15:04:05-0700"

var accessToken string

var eventsCache *cache.Cache

func init() {
	if !config.CheckHaveenv("FB_APP_ID", "FB_SECRET") {
		return
	}

	accessToken = getAccessToken()
	eventsCache = cache.New(1*time.Minute, 0)

	loadEvents()
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
func getEvents() ([]*Event, error) {

	req := requestWithData("GET",
		"https://graph.facebook.com/v2.12/hacksocmcr/events", map[string]string{
			"fields": "name,description,cover,place,start_time,end_time,attending_count,interested_count",
		}, true)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("status code %s instead of 2xx Success",
			response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

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
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
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

	return events, nil
}

type EventGroup struct {
	RightNow []*Event
	Upcoming []*Event
	Past     []*Event
}

func loadEvents() ([]*Event, error) {
	events, err := getEvents()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	eventsCache.Set("events", events, cache.DefaultExpiration)
	return events, nil
}
func getEventsFromCache() ([]*Event, error) {
	events, expiration, found := eventsCache.GetWithExpiration("events")

	if !found || !(time.Now().Before(expiration)) {
		var err error
		events, err = loadEvents()

		if err != nil {
			return nil, err
		}
	}
	return events.([]*Event), nil
}

// GroupEvents groups the events by 'right now', 'upcoming', and 'past'
func GroupEvents() (*EventGroup, error) {
	events, err := getEventsFromCache()

	if err != nil {
		return nil, err
	}

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
