package galleryService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type JSONPhotosets struct {
	Photosets struct {
		Photoset []struct {
			ID      string
			Primary string
			Secret  string
			Server  string
			Farm    int
			Title   struct {
				Content string
			}
			Description struct {
				Content string
			}
		}
	}
}

var jsonBytes []byte
var photosets JSONPhotosets

func init() {

	if os.Getenv("FLICKR_API_KEY") == "" {
		log.Println("Environment variable FLICKR_API_KEY is not assigned")
		return
	}

	if os.Getenv("FLICKR_USER_ID") == "" {
		log.Println("Environment variable FLICKR_USER_ID is not assigned")
		return
	}

	apiKey := os.Getenv("FLICKR_API_KEY")
	userID := os.Getenv("FLICKR_USER_ID")
	res, err := http.Get("https://api.flickr.com/services/rest/?&method=flickr.photosets.getList&api_key=" +
		apiKey + "&user_id=" + userID + "&format=json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	jsonBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	jsonBytes = jsonBytes[14 : len(jsonBytes)-1]
	jsonBytes = bytes.Replace(jsonBytes, []byte("_"), []byte(""), -1)

	err = json.Unmarshal(jsonBytes, &photosets)
	if err != nil {
		panic(err)
	}
}

func GetAlbums() JSONPhotosets {
	return photosets
}
