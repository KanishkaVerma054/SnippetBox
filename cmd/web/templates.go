package main

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"html/template"
	"path/filepath"
	"time"
)

/*
	// 5.6 Custom template functions

	// Create a humanDate function which returns a nicely formatted string
	// representation of a time.Time object.
*/
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

/*
	//  5.6 Custom template functions

	// Initialize a template.FuncMap object and store it in a global variable. This is
	// essentially a string-keyed map which acts as a lookup between the names of our
	// custom template functions and the functions themselves.
*/
var functions = template.FuncMap{
	"humanDate": humanDate,
}

/*
	// 5.1 Displaying dynamic data: Rendering multiple pieces of data

	// Define a templateData type to act as the holding structure for
	// any dynamic data that we want to pass to our HTML templates.
	// At the moment it only contains one field, but we'll add more
	// to it as the build progresses.

*/
type templateData struct {

	Snippet *models.Snippet

	/*
		// 5.2 Template actions and functions

		// Include a Snippets field in the templateData struct.
	*/
	Snippets []*models.Snippet

	/*	
		// 5.6 Common dynamic data

		// Add a CurrentYear field to the templateData struct.
	*/
	CurrentYear int
}

func newTemplateCache() (map[string]*template.Template, error) {

	/*
		// 5.3 Caching templates
		// Intialize a new map to act as the cache.
	*/
	cache := map[string]*template.Template{}

	/*
		// 5.3 Caching templates

		// Use the filepath.Glob() function to get a slice of all filepaths that
		// match the pattern "./ui/html/pages/*.tmpl". This will essentially gives
		// us a slice of all the filepaths for our application 'page' templates
		// like: [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]
	*/
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	/*
		// 5.3 Caching templates
		
		// Loop through the page filepaths one-by-one.
	*/
	for _, page := range pages {
		
		/*
			// 5.3 Caching templates
		
			// Extract the file name (like 'home.tmpl') from the full filepath
			// and assign it to the name variable.
		*/
		name := filepath.Base(page)

		/*
			// 5.3 Caching templates
		
			// Create a slice containing the filepaths for our base template, any
			// partials and the page.
		*/
		// files := []string{
		// 	"./ui/html/base.html",
		// 	"./ui/html/partials/nav.html",
		// 	page,
		// }

		/*
			// 5.3 Caching templates
		
			// Parse the files into a template set.
		*/
		// ts, err := template.ParseFiles(files...)
		// if err != nil {
		// 	return nil, err
		// }

		/*
			// 5.3 Caching templates

			// Parse the base template file into a template set.
		*/
		// ts, err := template.ParseFiles("./ui/html/base.html")
		// if err != nil {
		// 	return nil, err
		// }

		/*
			// 5.6 Custom template functions

			// The template.FuncMap must be registered with the template set before you
			// call the ParseFiles() method. This means we have to use template.New() to
			// create an empty template set, use the Funcs() method to register the
			// template.FuncMap, and then parse the file as normal.
		*/
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		/*
			// 5.3 Caching templates

			// Call ParseGlob() *on this template set* to add any partials.
		*/
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		/*
			// 5.3 Caching templates

			// Call ParseFiles() *on this template set* to add the page template.
		*/
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		/*
			// 5.3 Caching templates
		
			// Add the template set to the map, using the name of the page
			// (like 'home.html') as the key.
		*/
		cache[name] = ts
	}

	/*
		// 5.3 Caching templates
	
		// Return the map.
	*/
	return cache, nil
}