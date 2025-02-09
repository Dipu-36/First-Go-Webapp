package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//Appconfig holds the application config of the application to make it globally available
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}

