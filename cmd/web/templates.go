package main

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"KanishkaVerma054/snipperBox.dev/ui"
	"html/template"
	"io/fs"
	"path/filepath"
	"time"
)

func humanDate(t time.Time) string {
	// return t.Format("02 Jan 2006 at 15:04")

	/*
		// 14.1. Unit testing and sub-tests: Table-driven tests

		// Return the empty string if time has the zero value.
	*/
	if t.IsZero() {
		return ""
	}

	/*
		// 14.1. Unit testing and sub-tests: Table-driven tests

		// Convert the time to UTC before formatting it.
	*/
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	Snippet 		*models.Snippet
	Snippets 		[]*models.Snippet
	CurrentYear 	int
	Form			any
	Flash			string
	IsAuthenticated	bool
	CSRFToken		string
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	// pages, err := filepath.Glob("./ui/html/pages/*.html")
	// if err != nil {
	// 	return nil, err
	// }

	/*
		// 13.1 Using embedded files: Embedding HTML templates

		// Use fs.Glob() to get a slice of all filepaths in the ui.Files embedded
		// filesystem which match the pattern 'html/pages/*.tmpl'. This essentially
		// gives us a slice of all the 'page' templates for the application, just like before.
	*/
	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		
		name := filepath.Base(page)

		// ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		// if err != nil {
		// 	return nil, err
		// }

		// ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		// if err != nil {
		// 	return nil, err
		// }

		// ts, err = ts.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }

		/*
			// 13.1 Using embedded files: Embedding HTML templates

			// Create a slice containing the filepath patterns for the templates we
			// want to parse.
		*/
		patterns := []string {
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		/*
			// 13.1 Using embedded files: Embedding HTML templates

			// Use ParseFS() instead of ParseFiles() to parse the template files 
			// from the ui.Files embedded filesystem.
		*/
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}