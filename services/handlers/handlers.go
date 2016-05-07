package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"github.com/andreimuntean/www/services/minifier"
)

var templates map[string]*template.Template

// Execute loads the templates from the specified directory and configures the routes.
func Execute(templateDirectory string) error {
	sharedDirectory := filepath.Join(templateDirectory, "shared")
	
	// Loads a path for every template.
	templatePaths, err := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	
	if err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)		
	}
	
	// Loads a path for every shared template.
	sharedPaths, err := filepath.Glob(filepath.Join(sharedDirectory, "*.tmpl"))
	
	if err != nil {
		return fmt.Errorf("Could not find shared directory '%s'.", sharedDirectory)				
	}
	
	// Loads the templates.
	templates = make(map[string]*template.Template)
	
	for _, templatePath := range templatePaths {
		t, err := minifier.ParseTemplate(append(sharedPaths, templatePath)...)
		
		if err != nil {
			return err
		}
		
		templates[filepath.Base(templatePath)] = t
	}

	// Configures the routes.
	http.HandleFunc("/", indexHandler)

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates["index.tmpl"].ExecuteTemplate(w, "layout", nil)
}
