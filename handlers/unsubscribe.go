package handlers

import (
	"github.com/hacksoc-manchester/www/helpers/crypto"
	"github.com/hacksoc-manchester/www/services/databaseService"
	"net/http"
)

func unsubscribe(w http.ResponseWriter, r *http.Request) {
	email, err := crypto.Decrypt(r.FormValue("token"))

	if err != nil {
		templates["message"].ExecuteTemplate(w, "layout", messageModel{"Error", "Token is invalid."})
		return
	}

	if !databaseService.ExistsUser(email) {
		templates["message"].ExecuteTemplate(
			w, "layout", messageModel{"Error", `Email "` + email + `" is not part of our mailing list.`})
		return
	}

	if err := databaseService.DeleteUser(email); err != nil {
		templates["message"].ExecuteTemplate(w, "layout", messageModel{"Error", "An unexpected error has occurred. Please try again later."})
		return
	}

	templates["message"].ExecuteTemplate(
		w, "layout", messageModel{"Unsubscribed", `Email "` + email + `" has been removed from our mailing list.`})
}
