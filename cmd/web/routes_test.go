package main

import (
	"fmt"
	"testing"

	"github.com/andkolbe/bed-and-breakfast/internal/config"
	"github.com/go-chi/chi/v5"
)


func TestRoutes(t *testing.T) {
	var app config.AppConfig // need the app config to run our routes

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux: // mux is a pointer to chi.Mux so that is what we have to test for
		// do nothing; test passed
	default: 
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}