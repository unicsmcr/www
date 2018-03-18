package handlers

import (
	"net/http"

	"github.com/alexdmtr/www/services/galleryService"
)

func gallery(w http.ResponseWriter, r *http.Request) {
	albums := galleryService.GetAlbums()
	renderTemplate(w, r, "gallery", &albums)
}
