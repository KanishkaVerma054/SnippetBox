package main

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"strconv"
)

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}


	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n", snippet)
	// }
	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/pages/home.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err) 
	// 	return
	// }


	/*
		// 5.2 Template actions and functions

		// Create an instance of a templateData struct holding the slice of
		// snippets.
	*/
	// data := &templateData{
	// 	Snippets: snippets,
	// }

	/*
		// 5.2 Template actions and functions: Using the if and range actions

		// Pass in the templateData struct when executing the template.
	*/
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	/*
		// 5.3 Caching templates

		// Use the new render helper.
	*/
	// app.render(w, http.StatusOK, "home.html", &templateData{
	// 	Snippets: snippets,
	// })

	/*	
		// 5.5 Common dynamic data

		// Call the newTemplateData() helper to get a templateData struct containing
		// the 'default' data (which for now is just the current year), and add the
		// snippets slice to it.
	*/
	data := app.newTemplateData(r)
	data.Snippets = snippets

	/*	
		// 5.5 Common dynamic data

		// Pass the data to the render() helper as normal.
	*/
	app.render(w, http.StatusOK, "home.html", data)
}

func(app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) 
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// fmt.Fprintf(w, "%+v", snippet)
	/*
		5.1 Displaying dynamic data

		// Initialize a slice containing the paths to the view.tmpl file,
		// plus the base layout and navigation partial that we made earlier.
	*/
	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/pages/view.html",
	// }

	/*
		// 5.1 Displaying dynamic data

		// Parse the template files...
	*/
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	/*
		5.1 Displaying dynamic data: Rendering multiple pieces of data

		// Create an instance of a templateData struct holding the snippet data.
	*/
	// data := &templateData{
	// 	Snippet: snippet,
	// }

	
	/*
		// 5.1 Displaying dynamic data

		// And then execute them. Notice how we are passing in the snippet
		// data (a models.Snippet struct) as the final parameter?
	*/
	// err = ts.ExecuteTemplate(w, "base", snippet)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	/*
		5.1 Displaying dynamic data: Rendering multiple pieces of data

		// Pass in the templateData struct when executing the template.
	*/
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }	

	/*
		// 5.3 Caching Templates

		// Use the new render helper.
	*/
	// app.render(w, http.StatusOK, "view.html", &templateData{
	// 	Snippet: snippet,
	// })

	/*	
		// 5.5 Common dynamic data

		// And do the same thing again here...
	*/
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.html", data)

}

func(app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji, \nBut slowly, slowly!\n\n-Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}