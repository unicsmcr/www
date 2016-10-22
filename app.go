package main

import (
	"github.com/tdewolff/minify"
	"github.com/hacksoc-manchester/www/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Makes the assets folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Sets up the routes.
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")

	if err := handlers.Execute(templateDirectory); err != nil {
		log.Fatal(err)
	}

	// Starts the server.
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
