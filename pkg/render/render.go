package render

import (
	"bytes"
	"github.com/borntodie-new/information/pkg/config"
	"github.com/borntodie-new/information/pkg/model"

	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// create a template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
		app.UseCache = true
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cacheTemplate := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./template/*.page.tmpl")
	if err != nil {
		return cacheTemplate, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cacheTemplate, err
		}
		// parse all the file named *.layout.tmpl from ./templates
		matches, err := filepath.Glob("./template/*.layout.tmpl")
		if err != nil {
			return cacheTemplate, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.tmpl")
			if err != nil {
				return cacheTemplate, err
			}
		}
		cacheTemplate[name] = ts
	}
	return cacheTemplate, nil
}
