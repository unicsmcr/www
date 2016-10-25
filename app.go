package main

import (
<<<<<<< HEAD
=======
	"github.com/hacksoc-manchester/www/handlers"
>>>>>>> 98b918ba286ba79143133744c7c913ba3c2b7bdd
	"log"
	"net/http"
	"os"
	"path/filepath"
<<<<<<< HEAD

	"github.com/hacksoc-manchester/www/handlers"
=======
>>>>>>> 98b918ba286ba79143133744c7c913ba3c2b7bdd
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
