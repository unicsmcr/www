package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hacksoc-manchester/www/handlers"
)

func main() {
	// Make the "assets" folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up the routes.
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")

	if err := handlers.Execute(templateDirectory); err != nil {
		log.Fatal(err)
	}

	// Start the server.
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
