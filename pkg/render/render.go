package render

import (
	"github.com/andkolbe/bed-and-breakfast/pkg/config"
	"github.com/andkolbe/bed-and-breakfast/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// a FuncMap is a map of functions that can be used in a template
// Go allows us to create our own functions that aren't included in Go, and pass them to the templates
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// add data from the templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// in dev mode, don't use the template cache, instead rebuild it on every request
	var tc map[string]*template.Template

	if app.UseCache { // if use cache is true, use the template cache
		tc = app.TemplateCache
	} else { // else, rebuild a new template cache on every request
		tc, _ = CreateTemplateCache()
	}

	// create template cache when the app starts, then when we render a page, we are pulling a value from our config
	// get the template cache from the app config
	// tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// writing to a buffer lets me check to see if there's an error and determine where it comes from more easily
	// if you write straight to the response writer, you can still check for an error, but have to do some extra work to see where it comes from 
	// a bytes buffer will hold the parsed template in bytes in memory
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td) // take the template we have, execute it, don't pass it any data and store the value in the buffer variable

	// writes from the buffer to the response writer
	_, err := buf.WriteTo(w) // returns the number of bytes but we don't care about that so use _
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// create a template cache that holds all our html templates in a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// we use a map because it lets us looks up values very quickly
	myCache := map[string]*template.Template{}

	// go to the templates folder, and get all of the files that start with anything but end with .page.html
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// loop through all of our templates that end in page.html
	for _, page := range pages { // pul out the names of the pages
		name := filepath.Base(page)
		// ts = template set
		// create a template set based upon the loop of pages and pass it our created functions
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// does this template match any layouts? Should we use a layout that's defined for this template with this particular template
		// look for the existance of any layouts in this particular folder called templates

		// go to the templates folder, and get all of the files that end with .layout.html
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if a .layout.html match is found, the length will be greater than 0
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// take the template set we created and add it to the cache
		myCache[name] = ts
	}

	return myCache, nil
}
