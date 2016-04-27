package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/hacksoc-manchester/www/handlers"
)

var assetDirectory string
var templateDirectory string

func init() {
	dir, _ := os.Getwd()
	assetDirectory = filepath.Join(dir, "assets")
	templateDirectory = filepath.Join(dir, "templates")
}

func main() {
	if err := handlers.Execute(assetDirectory, templateDirectory); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
