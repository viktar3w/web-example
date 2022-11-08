package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// SessionLoad load and save every session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// NoSurf add CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
