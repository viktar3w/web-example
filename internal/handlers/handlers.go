package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/viktar3w/web-example/internal/config"
	"github.com/viktar3w/web-example/internal/models"
	"github.com/viktar3w/web-example/internal/render"
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
	if err := render.RenderTemplate(w, r, "home.page.tmpl", data); err != nil {
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

	if err := render.RenderTemplate(w, r, "about.page.tmpl", data); err != nil {
		log.Println("Error at About Page: %s\n", err)
	}
}

// Generals ...
func (rep *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Generals Page",
		},
	}
	if err := render.RenderTemplate(w, r, "generals.page.tmpl", data); err != nil {
		log.Println("Error at Generals Page: %s\n", err)
	}
}

// Majors ...
func (rep *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Major's Suite Page",
		},
	}
	if err := render.RenderTemplate(w, r, "majors.page.tmpl", data); err != nil {
		log.Println("Error at Majors Page: %s\n", err)
	}
}

// Reservation ...
func (rep *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Reservation Page",
		},
	}
	if err := render.RenderTemplate(w, r, "make-reservation.page.tmpl", data); err != nil {
		log.Println("Error at Reservation Page: %s\n", err)
	}
}

// Availability ...
func (rep *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Reservation Page",
			"start": r.URL.Query().Get("start"),
			"end":   r.URL.Query().Get("end"),
		},
	}
	if err := render.RenderTemplate(w, r, "search-availability.page.tmpl", data); err != nil {
		log.Println("Error at Availability Page: %s\n", err)
	}
}

// PostAvailability ...
func (rep *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("Post to search Availability start at `%s` and end `%s`", start, end)))
	if err != nil {
		log.Println(err)
		return
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson ...
func (rep *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
		return
	}
}

// Contact ...
func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	data := &models.TemplateData{
		StringMap: map[string]string{
			"title": "Contact Page",
		},
	}
	if err := render.RenderTemplate(w, r, "contact.page.tmpl", data); err != nil {
		log.Println("Error at Contact Page: %s\n", err)
	}
}
