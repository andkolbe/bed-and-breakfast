package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/andkolbe/bed-and-breakfast/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager // store the session in here so it can be used in the main and handler packages
	MailChan      chan models.MailData
}
