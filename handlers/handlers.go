package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates map[string]*template.Template

// Execute loads templates from the specified directory and configures routes.
func Execute(templateDirectory string) error {
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)		
	}
	
	// Loads a path for every template.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	
	// Loads a path for every shared template.
	sharedPaths, _ := filepath.Glob(filepath.Join(templateDirectory, "shared/*.tmpl"))
	
	// Loads the templates.
	templates = make(map[string]*template.Template)
	
	for _, templatePath := range templatePaths {
		t, err := template.ParseFiles(append(sharedPaths, templatePath)...)
		
		if err != nil {
			return err
		}
		
		name := strings.Split(filepath.Base(templatePath), ".")[0]
		templates[name] = t
	}

	// Configures the routes.
	http.HandleFunc("/", indexHandler)

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates["index"].ExecuteTemplate(w, "layout", nil)
}
