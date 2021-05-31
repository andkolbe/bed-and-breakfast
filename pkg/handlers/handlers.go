package handlers

import (
	"net/http"

	"github.com/andkolbe/bed-and-breakfast/pkg/config"
	"github.com/andkolbe/bed-and-breakfast/pkg/models"
	"github.com/andkolbe/bed-and-breakfast/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.html", &models.TemplateData{})
}

// room1 page handler
func (m *Repository) Room1(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "room1.page.tmpl", &models.TemplateData{})
}

// room2 page handler
func (m *Repository) Room2(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "room2.page.tmpl", &models.TemplateData{})
}
