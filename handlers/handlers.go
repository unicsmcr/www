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
	
	// Loads template paths.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
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
	http.HandleFunc("/events", eventsHandler)

	return nil
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	templates["events"].ExecuteTemplate(w, "layout", nil)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	var args struct {
		HasUpcomingEvent bool
		EventName string
		EventDate string
		Menu []struct {
			Name string
			Tag string
		}
	}
	
	args.HasUpcomingEvent = true
	args.EventName = "GreatUniHack 2016"
	args.EventDate = "September 16 at 5:30pm"
	args.Menu = []struct {
		Name string
		Tag string
	}{
		{"Events", "events"},
		{"Team", "team"},
		{"Gallery", "gallery"},
		{"Contact", "contact"},
	}
  	
	templates["index"].ExecuteTemplate(w, "layout", &args)
}
