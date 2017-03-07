package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type JSONInfo struct {
	Content string
}

type JSONPhoto struct {
	ID                  string
	Primary             string
	Secret              string
	Server              string
	Farm                int
	Photos              string
	Videos              int
	Title               JSONInfo
	Description         JSONInfo
	NeedsInterstitial   int
	VisibilityCanSeeSet int
	CountViews          string
	CountComments       string
	CanComment          int
	CateCreate          string
	CateUpdate          string
}

type JSONObject struct {
	Photosets struct {
		Page     int
		Pages    int
		Perpage  int
		Potal    int
		Photoset []JSONPhoto
	}
}

var albums JSONObject

func init() {
	path, _ := filepath.Abs("assets/json/flickr.json")
	file, e := ioutil.ReadFile(path)

	if e != nil {
		log.Fatal(e)
	}

	json.Unmarshal(file, &albums)

}

func gallery(w http.ResponseWriter, r *http.Request) {
	templates["gallery"].ExecuteTemplate(w, "layout", &albums)
}
