package render

import (
	"bytes"
	"html/template"
	"log"
	"github.com/Dipu-36/Go-webapp/pkg/config"
	"github.com/Dipu-36/Go-webapp/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a 
}

func AddDefaultData ( td *models.TemplateData) *models.TemplateData {

	return td
}
// RenderTemplates renders a template
func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from the app config

	//We are using the global variable to access and store the already cached templates inside a new variable 
	var templateCache map[string]*template.Template
	
	if app.UseCache{
		//get the template cache from the app config
		templateCache = app.TemplateCache
	}else {
		templateCache, _ = CreateTemplateCache()
	}
	
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not find template in cache")
	}

	// Render the template to a buffer
	buff := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buff, td)

	// Write the rendered template to the response writer
	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println("Error writing to response:", err)
	}
}

// createTemplateCache creates a template cache from all page templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	//Glob returns a slice of string 
	if err != nil {
		return myCache, err
	}

	// Range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)

		// Parse each page template file
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		

		// Add the parsed template to the cache
		myCache[name] = templateSet
	}

	return myCache, nil
}
