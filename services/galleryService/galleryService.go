package galleryService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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

type Album struct {
	Title       string
	Description string
	Link        string
	ImageURL    string
}

var jsonBytes []byte
var photos JSONPhotosets

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
	hc := http.Client{}
	apiURL := "https://api.flickr.com/services/rest/"
	form := url.Values{}
	form.Add("method", "flickr.photosets.getList")
	form.Add("format", "json")
	form.Add("api_key", apiKey)
	form.Add("user_id", userID)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := hc.Do(req)
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

	err = json.Unmarshal(jsonBytes, &photos)
	if err != nil {
		panic(err)
	}
}

func GetAlbums() []map[string]string {

	var albums []map[string]string
	for _, photoset := range photos.Photosets.Photoset {

		album := map[string]string{
			"Title":       photoset.Title.Content,
			"Description": photoset.Description.Content,
			"Link":        "http://www.flickr.com/photos/" + os.Getenv("FLICKR_USER_ID") + "/sets/" + photoset.ID,
			"ImageURL":    "https://farm" + strconv.Itoa(photoset.Farm) + ".staticflickr.com/" + photoset.Server + "/" + photoset.Primary + "_" + photoset.Secret + ".jpg",
		}
		albums = append(albums, album)
	}
	return albums
}
