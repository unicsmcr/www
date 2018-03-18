package handlers

import (
	"net/http"

	"github.com/hacksoc-manchester/www/helpers/crypto"
	"github.com/hacksoc-manchester/www/services/databaseService"
)

func unsubscribe(w http.ResponseWriter, r *http.Request) {
	handleUserError := func(message messageModel) {
		renderTemplate(w, r, "message", message)
	}
	email, err := crypto.Decrypt(r.FormValue("token"))

	if err != nil {
		handleUserError(messageModel{"Error", "Token is invalid."})
		return
	}

	if !databaseService.ExistsUser(email) {
		handleUserError(messageModel{"Error", `Email "` + email + `" is not part of our mailing list.`})
		return
	}

	if err := databaseService.DeleteUser(email); err != nil {
		handleUserError(messageModel{"Error", "An unexpected error has occurred. Please try again later."})
		return
	}

	renderTemplate(
		w, r, "message", messageModel{"Unsubscribed", `Email "` + email + `" has been removed from our mailing list.`})
}
