package galleryService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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

var albums []map[string]string

const flickrApiURL = "https://api.flickr.com/services/rest/"

func init() {
	if os.Getenv("FLICKR_API_KEY") == "" {
		log.Println("Environment variable FLICKR_API_KEY is not assigned.")
		return
	}

	if os.Getenv("FLICKR_USER_ID") == "" {
		log.Println("Environment variable FLICKR_USER_ID is not assigned.")
		return
	}

	flickrApiKey := os.Getenv("FLICKR_API_KEY")
	flickrUserID := os.Getenv("FLICKR_USER_ID")
	httpClient := http.Client{}
	form := url.Values{}
	form.Add("method", "flickr.photosets.getList")
	form.Add("format", "json")
	form.Add("api_key", flickrApiKey)
	form.Add("user_id", flickrUserID)

	req, err := http.NewRequest("POST", flickrApiURL, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var jsonBytes []byte
	jsonBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// The first 14 characters and the last one are always "jsonFlickrApi(" and ")"
	// They need to be removed, otherwise the json won't unmarshal
	jsonBytes = jsonBytes[14 : len(jsonBytes)-1]
	// The titles and descriptions of the albums returned are stored in fields called
	// "_content". The underscore makes the unmarshal function ignore them, so they have
	// to be removed.
	jsonBytes = bytes.Replace(jsonBytes, []byte("_"), []byte(""), -1)

	var photos JSONPhotosets
	if err = json.Unmarshal(jsonBytes, &photos); err != nil {
		panic(err)
	}

	for _, photoset := range photos.Photosets.Photoset {

		albumLink := fmt.Sprintf("http://www.flickr.com/photos/%s/sets/%s",
			os.Getenv("FLICKR_USER_ID"), photoset.ID)
		albumImage := fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s.jpg",
			photoset.Farm, photoset.Server, photoset.Primary, photoset.Secret)
		album := map[string]string{
			"Title":       photoset.Title.Content,
			"Description": photoset.Description.Content,
			"Link":        albumLink,
			"ImageURL":    albumImage,
		}
		albums = append(albums, album)
	}
}

func GetAlbums() []map[string]string {
	return albums
}
