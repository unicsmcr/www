package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/hacksoc-manchester/www/config"
	"github.com/hacksoc-manchester/www/handlers"
)

func main() {
	// Make the "assets" folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up the routes.
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")

	err := handlers.Execute(templateDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Start the server.
	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		err = http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		err = http.ListenAndServe(":8080", nil)
	}

	if err != nil {
		log.Fatal(err)
	}

}
