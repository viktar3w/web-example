package render

import (
	"bytes"
	"errors"
	"github.com/viktar3w/web-example/pkg/config"
	"github.com/viktar3w/web-example/pkg/models"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	baseTemplatePath = "./templates/"
)

var (
	functions = template.FuncMap{}
	app       *config.AppConfig
)

// NewTemplates ...
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, templatePath string, td *models.TemplateData) error {
	var templates map[string]*template.Template
	var err error
	if app.UseCache {
		templates = app.TemplateCache
	} else {
		templates, err = CreateTemplateCache()
		if err != nil {
			return err
		}
	}
	parsedTemplate, ok := templates[templatePath]
	if !ok {
		return errors.New("Template isn't - " + templatePath)
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	if err = parsedTemplate.Execute(buf, td); err != nil {
		return err
	}
	_, err = buf.WriteTo(w)
	return err
}

// CreateTemplateCache ...
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(baseTemplatePath + "*.page.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob(baseTemplatePath + "*.layout.tmpl")
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(baseTemplatePath + "*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
