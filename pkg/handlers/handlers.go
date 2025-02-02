package handlers

import (
	"github.com/Dipu-36/Go-webapp/pkg/config"
	"github.com/Dipu-36/Go-webapp/pkg/models"

	"github.com/Dipu-36/Go-webapp/pkg/render"
	"net/http"
)

//Repo the repository used by handlers
var Repo *Repository

//Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

func NewRepo(a * config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}
//Home is a homepage handler function
func ( m *Repository ) Home(w http.ResponseWriter, r *http.Request){
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{} )
}
// About is the about page Handler
func ( m *Repository ) About(w http.ResponseWriter, r *http.Request){
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	//send the data to the template
	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{
		StringMap : stringMap,
	})
}

