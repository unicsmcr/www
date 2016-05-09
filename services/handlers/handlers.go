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

// Execute loads the templates from the specified directory and configures the routes.
func Execute(templateDirectory string) error {
	sharedDirectory := filepath.Join(templateDirectory, "shared")
	
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)		
	}
	
	if _, err := os.Stat(sharedDirectory); err != nil {
		return fmt.Errorf("Could not find shared directory '%s'.", sharedDirectory)				
	}
	
	// Loads a path for every template.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	
	// Loads a path for every shared template.
	sharedPaths, _ := filepath.Glob(filepath.Join(sharedDirectory, "*.tmpl"))
	
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
	
	// Makes the assets folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Configures the routes.
	http.HandleFunc("/", indexHandler)

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates["index"].ExecuteTemplate(w, "layout", nil)
}
