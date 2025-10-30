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

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

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

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)


	// }
}

func(app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) 
		return
	}

	/*
		Single-record SQL queries: Using the model in our handlers

		// Use the SnippetModel object's Get method to retrieve the data for a
		// specific record based on its ID. If no matching record is found,
		// return a 404 Not Found response.
	*/
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

	/*
		// 4.7 Single-record SQL queries: Using the model in our handlers

		// Write the snippet data as a plain-text HTTP response body.
	*/
	fmt.Fprintf(w, "%+v", snippet)
}

func(app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}

	/*
		// 4.6 Executing SQL statements

		// Creatimg some variables holding dummy data.
	*/
	title := "O snail"
	content := "O snail\nClimb Mount Fuji, \nBut slowly, slowly!\n\n-Kobayashi Issa"
	expires := 7

	/*
		// 4.6 Executing SQL statements

		// Pass the data to the SnippetModel.Insert() method, receiving the
		// ID of the new record back.
	*/
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	/*
		// 4.6 Executing SQL statements

		// Redirect the user to the relevant page for the snippet.
	*/
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}