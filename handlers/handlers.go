package handlers

import (
	"fmt"
	"github.com/haisum/recaptcha"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type messageModel struct {
	Title   string
	Message string
}

var templates map[string]*template.Template
var reCaptchaSiteKey = os.Getenv("RECAPTCHA_SITE_KEY")
var reCaptcha = recaptcha.R{
	Secret: os.Getenv("RECAPTCHA_SECRET_KEY"),
}

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
	http.HandleFunc("/", index)
	http.HandleFunc("/events", events)
	http.HandleFunc("/team", team)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/partners", partners)
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/unsubscribe", unsubscribe)

	return nil
}
