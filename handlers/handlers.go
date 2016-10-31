package handlers

import (
	"fmt"
	"github.com/haisum/recaptcha"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"html/template"
	"io/ioutil"
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

var m = minify.New()

func compileTemplates(templatePaths ...string) (*template.Template, error) {
	var tmpl *template.Template
	for _, templatePath := range templatePaths {
		name := filepath.Base(templatePath)
		if tmpl == nil {
			tmpl = template.New(name)
		} else {
			tmpl = tmpl.New(name)
		}
		b, err := ioutil.ReadFile(templatePath)
		if err != nil {
			return nil, err
		}

		mb, err := m.Bytes("text/html", b)
		if err != nil {
			return nil, err
		}
		tmpl.Parse(string(mb))
	}
	return tmpl, nil
}

func minifyCSSFiles(templateDirectory string) {
	cssFileDirectory := filepath.Join(templateDirectory, "../assets/css/")
	cssFilePaths, _ := filepath.Glob(filepath.Join(cssFileDirectory, "*.css"))
	for _, cssFilePath := range cssFilePaths {
		if strings.HasSuffix(cssFilePath, ".min.css") {
			continue
		}
		cssFile, err := ioutil.ReadFile(cssFilePath)
		if err != nil {
			panic(err)
		}
		cssFile, err = m.Bytes("text/css", cssFile)
		if err != nil {
			panic(err)
		}
		minCSSFilePath := strings.Replace(cssFilePath, ".css", ".min.css", 1)
		err = ioutil.WriteFile(minCSSFilePath, cssFile, 0666)
		if err != nil {
			panic(err)
		}
	}
}

// Execute loads templates from the specified directory and configures routes.
func Execute(templateDirectory string) error {
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)
	}

	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	minifyCSSFiles(templateDirectory)

	// Loads template paths.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	sharedPaths, _ := filepath.Glob(filepath.Join(templateDirectory, "shared/*.tmpl"))

	// Loads the templates.
	templates = make(map[string]*template.Template)
	for _, templatePath := range templatePaths {
		tmpl := template.Must(compileTemplates(append(sharedPaths, templatePath)...))

		name := strings.Split(filepath.Base(templatePath), ".")[0]
		templates[name] = tmpl
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
