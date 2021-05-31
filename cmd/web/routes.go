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
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/room1", handlers.Repo.Room1)
	mux.Get("/room2", handlers.Repo.Room2)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)	
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)


	// create a file server - a place for the routes to get static files from. Must do this for Go tmpls
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

// each request to a url is stateless, what can we do to make sure our user data is saved on each request? Sessions
