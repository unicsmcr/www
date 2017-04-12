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
)

// An Album represents a HackSoc album.
type Album struct {
	Title       string
	Description string
	Link        string
	ImageURL    string
}

var albums []map[string]string

const flickrAPIURL = "https://api.flickr.com/services/rest"

func init() {
	if os.Getenv("FLICKR_API_KEY") == "" {
		log.Println("Environment variable FLICKR_API_KEY is not assigned.")
		return
	}

	if os.Getenv("FLICKR_USER_ID") == "" {
		log.Println("Environment variable FLICKR_USER_ID is not assigned.")
		return
	}

	form := url.Values{}
	form.Add("method", "flickr.photosets.getList")
	form.Add("format", "json")
	form.Add("nojsoncallback", "true")
	form.Add("api_key", os.Getenv("FLICKR_API_KEY"))
	form.Add("user_id", os.Getenv("FLICKR_USER_ID"))

	req, _ := http.NewRequest("POST", flickrAPIURL, bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = form

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	jsonBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var photos struct {
		Photosets struct {
			Photoset []struct {
				ID      string
				Primary string
				Secret  string
				Server  string
				Farm    int
				Title   struct {
					Content string `json:"_content"`
				}
				Description struct {
					Content string `json:"_content"`
				}
			}
		}
	}

	if err = json.Unmarshal(jsonBytes, &photos); err != nil {
		panic(err)
	}

	for _, photoset := range photos.Photosets.Photoset {
		// Get the album link.
		link := fmt.Sprintf("https://www.flickr.com/photos/%s/sets/%s",
			os.Getenv("FLICKR_USER_ID"), photoset.ID)

		// Get the album thumbnail URL.
		thumbnailURL := fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s.jpg",
			photoset.Farm, photoset.Server, photoset.Primary, photoset.Secret)

		album := map[string]string{
			"Title":       photoset.Title.Content,
			"Description": photoset.Description.Content,
			"Link":        link,
			"ImageURL":    thumbnailURL,
		}

		albums = append(albums, album)
	}
}

// GetAlbums gets the albums.
func GetAlbums() []map[string]string {
	return albums
}
