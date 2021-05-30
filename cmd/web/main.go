package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/andkolbe/bed-and-breakfast/pkg/config"
	"github.com/andkolbe/bed-and-breakfast/pkg/handlers"
	"github.com/andkolbe/bed-and-breakfast/pkg/render"
)

const portNumber = ":8080"

// put these outside of the main function so they can be used everywhere in the package main
var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // true = cookie persists even if the browser window closes
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Staring application on port %s", portNumber)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
