package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"

	"github.com/haisum/recaptcha"
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

//the function for the minifying
var m = minify.New()

func compileTemplates(filenames ...string) (*template.Template, error) {

	var tmpl *template.Template
	for _, filename := range filenames {
		name := filepath.Base(filename)
		if tmpl == nil {
			tmpl = template.New(name)
		} else {
			tmpl = tmpl.New(name)
		}

		b, err := ioutil.ReadFile(filename)
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
		if cssFilePath[len(cssFilePath)-8:len(cssFilePath)] != ".min.css" {
			cssFile, err := ioutil.ReadFile(cssFilePath)
			if err != nil {
				panic(err)
			}
			cssFile, err = m.Bytes("text/css", cssFile)
			if err != nil {
				panic(err)
			}
			cssFilePathBase := filepath.Base(cssFilePath)
			miniCSSFilePath := filepath.Join(cssFileDirectory, cssFilePathBase[0:len(cssFilePathBase)-3]) + "min.css"
			err = ioutil.WriteFile(miniCSSFilePath, cssFile, 0666)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Execute loads templates from the specified directory and configures routes.
func Execute(templateDirectory string) error {
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)
	}

	//minify the css files
	m.AddFunc("text/css", css.Minify)
	minifyCSSFiles(templateDirectory)

	// Loads template paths.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	sharedPaths, _ := filepath.Glob(filepath.Join(templateDirectory, "shared/*.tmpl"))

	// Loads the templates.
	templates = make(map[string]*template.Template)

	m.AddFunc("text/html", html.Minify)

	for _, templatePath := range templatePaths {
		t := template.Must(compileTemplates(append(sharedPaths, templatePath)...))

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
