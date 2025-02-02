package main

import (
	"fmt"
	"log"
	"github.com/Dipu-36/Go-webapp/pkg/config"
	"github.com/Dipu-36/Go-webapp/pkg/handlers"
	"github.com/Dipu-36/Go-webapp/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
//main  holds the main application function
func main(){
	
	//change this to true when in production
	app.InProduction = false

	session = scs.New() 
	session.Lifetime = 24 *time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatalf("cannot create template cache: %v", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false
	 
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	
	srv := &http.Server {
		Addr : portNumber,
		Handler : routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
}