package main

import (
	"github.com/hacksoc-manchester/www/handlers"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadJSONfiles() {
	apiKey := "2346b05672267fe26f4635e7958816eb"
	userID := "147757828@N02"
	res, err := http.Get("https://api.flickr.com/services/rest/?&method=flickr.photosets.getList&api_key=" +
		apiKey + "&user_id=" + userID + "&format=json")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Create("./assets/json/flickr.json")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(jsonFile, res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	jsonFile.Close()
	jsonBytes, err := ioutil.ReadFile("assets/json/flickr.json")
	if err != nil {
		panic(err)
	}
	// The Flickr api returns a json file that starts with jsonFlickrApi( and ends with )
	// I need to remove these two parts in order for the json file to be read correctly by javascript
	err = ioutil.WriteFile("assets/json/flickr.json", jsonBytes[14:len(jsonBytes)-1], 0666)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Makes the assets folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Sets up the routes.
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")

	if err := handlers.Execute(templateDirectory); err != nil {
		log.Fatal(err)
	}

	// Download the Flickr JSON files for the Gallery
	downloadJSONfiles()

	// Starts the server.
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
