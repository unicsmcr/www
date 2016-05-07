package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/hacksoc-manchester/www/services/handlers"
)

func main() {
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")
	
	if err := handlers.Execute(templateDirectory); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
