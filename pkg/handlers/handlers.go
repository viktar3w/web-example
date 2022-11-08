package handlers

import (
	"github.com/viktar3w/web-example/pkg/config"
	"github.com/viktar3w/web-example/pkg/models"
	"github.com/viktar3w/web-example/pkg/render"
	"log"
	"net/http"
)

var (
	Repo *Repository
)

type Repository struct {
	App *config.AppConfig
}

// NewRepository ...
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers ...
func NewHandlers(r *Repository) {
	Repo = r
}

// Home ...
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rep.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Home Page",
		},
	}
	if err := render.RenderTemplate(w, "home.page.tmpl", data); err != nil {
		log.Println("Error at Home Page: %s\n", err)
	}
}

// About ...
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title":     "About Page",
			"remote_ip": rep.App.Session.GetString(r.Context(), "remote_ip"),
		},
	}

	if err := render.RenderTemplate(w, "about.page.tmpl", data); err != nil {
		log.Println("Error at About Page: %s\n", err)
	}
}
