package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var sponsors []struct {
	ID string
	URL string
}

func init() {
	path, _ := filepath.Abs("assets/json/sponsors.json")
	file, e := ioutil.ReadFile(path)
	
	if e != nil {
		log.Fatal(e)
	}
	
	json.Unmarshal(file, &sponsors)
}

func partners(w http.ResponseWriter, r *http.Request) {
	templates["partners"].ExecuteTemplate(w, "layout", &sponsors)
}