package minifier

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var m *minify.M

func init()  {
    m = minify.New()
    m.AddFunc("text/css", css.Minify)
    m.AddFunc("text/html", html.Minify)
    m.AddFunc("text/javascript", js.Minify)
}

// Parse loads and minifies the specified file.
func Parse(mediaType string, path string) (string, error) {
    file, err := ioutil.ReadFile(path)
		
    if err != nil {
        return "", err
    }
        
    minifiedFile, err := m.Bytes(mediaType, file)
    
    return string(minifiedFile), err
}

// ParseTemplate acts like template.ParseFiles but minifies the files.
func ParseTemplate(paths ...string) (*template.Template, error) {
	var t *template.Template
	
	for _, path := range paths {
		html, err := Parse("text/html", path)
		
		if err != nil {
			return nil, err
		}
		
		if (t == nil) {
			t = template.New(filepath.Base(path))
		} else {
			t = t.New(filepath.Base(path))
		}
		
		t.Parse(html)
	}
	
	return t, nil
}