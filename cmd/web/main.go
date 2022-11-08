package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/viktar3w/web-example/pkg/config"
	"github.com/viktar3w/web-example/pkg/handlers"
	"github.com/viktar3w/web-example/pkg/render"
	"log"
	"net/http"
	"time"
)

const (
	add  = ""
	port = 32121
	syf  = ""
	lang = "en"
)

var (
	app     = config.AppConfig{}
	session *scs.SessionManager
)

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = time.Hour + 3
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}
	app.TemplateCache = templateCache
	app.UseCache = app.InProduction
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	log.Println("Starting App")
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", add, port),
		Handler: routes(&app),
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
