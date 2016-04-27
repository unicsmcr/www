package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var assetDirectory string
var templateDirectory string

func Execute(assets string, templates string) error {
	assetDirectory = assets
	templateDirectory = templates

	if _, err := os.Stat(assetDirectory); err != nil {
		return fmt.Errorf("Could not find asset directory '%s'.", assetDirectory)
	}

	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)
	}

	// Configures the routes.
	http.HandleFunc("/", indexHandler)

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join(templateDirectory, "index.tmpl")

	if t, err := template.ParseFiles(templatePath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t.Execute(w, nil)
	}
}
