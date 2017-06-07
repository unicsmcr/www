package handlers

import (
	"fmt"
	"net/http"
)

func robots(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User-agent: *")
	fmt.Fprintln(w, "Disallow: /assets/")
}
