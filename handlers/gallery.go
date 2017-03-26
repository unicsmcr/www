package handlers

import (
	"github.com/hacksoc-manchester/www/services/galleryService"
	"net/http"
)

var albums galleryService.JSONPhotosets

func init() {
	albums = galleryService.GetAlbums()
}

func gallery(w http.ResponseWriter, r *http.Request) {
	templates["gallery"].ExecuteTemplate(w, "layout", &albums)
}
