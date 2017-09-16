package handlers

import "net/http"

func qr(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.facebook.com/events/140667106536809/", 302)
}
