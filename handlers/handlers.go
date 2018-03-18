package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/haisum/recaptcha"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"

	"github.com/oxtoacart/bpool"
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

var bufpool *bpool.BufferPool

var m = minify.New()

func init() {
	bufpool = bpool.NewBufferPool(64)
}
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

// renderTemplate is a wrapper around template.ExecuteTemplate.
// It writes into a bytes.Buffer before writing to the http.ResponseWriter to catch
// any errors resulting from populating the template.
// It also includes custom error handling.
func renderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) error {
	handleError := func(err error) {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
	}
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		err := fmt.Errorf("the template %s does not exist ", name)
		handleError(err)
		return err
	}

	// Create a buffer to temporarily write to and check if any errors were encounted.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, "layout", data)
	if err != nil {
		handleError(err)
		return err
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}

// Execute loads templates from the specified directory and configures routes.
func Execute(templateDirectory string) error {
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'", templateDirectory)
	}

	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	minifyCSSFiles(templateDirectory)

	// Load template paths.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	sharedPaths, _ := filepath.Glob(filepath.Join(templateDirectory, "shared/*.tmpl"))

	// Load the templates.
	templates = make(map[string]*template.Template)
	for _, templatePath := range templatePaths {
		tmpl := template.Must(compileTemplates(append(sharedPaths, templatePath)...))

		name := strings.Split(filepath.Base(templatePath), ".")[0]
		templates[name] = tmpl
	}

	// Configure the routes.
	http.HandleFunc("/", index)
	http.HandleFunc("/robots.txt", robots)
	http.HandleFunc("/events", events)
	http.HandleFunc("/team", team)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/partners", partners)
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/unsubscribe", unsubscribe)

	return nil
}
