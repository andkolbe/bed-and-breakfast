package main

import (
	"net/http"

	"github.com/andkolbe/bed-and-breakfast/pkg/config"
	"github.com/andkolbe/bed-and-breakfast/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// mux - multiplexer. a http handler that we create
	mux := chi.NewRouter()

	// middleware allows you to process a request as it comes in to the web app and perform some action on it
	// all of the routes go thru the middleware before running
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

// each request to a url is stateless, what can we do to make sure our user data is saved on each request? Sessions
