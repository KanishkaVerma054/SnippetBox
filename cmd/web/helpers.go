package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func(app *application) serverError(w http.ResponseWriter, err error){

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func(app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func(app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func(app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {

	/*
		// 5.3 Caching Templates

		// Retrieve the appropriate template set from the cache based on the page
		// name (like 'home.tmpl'). If no entry exists in the cache with the
		// provided name, then create a new error and call the serverError() helper
		// method that we made earlier and return.	
	*/
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	/*
		// 5.4 Catching runtime errors

		// Initialize a new buffer.
	*/
	buf := new(bytes.Buffer)

	/*
		// 5.4 Catching runtime errors

		// Write the template to the buffer, instead of straight to the
		// http.ResponseWriter. If there's an error, call our serverError() helper
		// and then return.
	*/
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	/*
		// 5.3 Caching Templates

		// Write out the provided HTTP status code ('200 OK', '400 Bad Request' // etc).
	*/
	/*
		// 5.4 Catching runtime errors

		// If the template is written to the buffer without any errors, we are safe
		// to go ahead and write the HTTP status code to http.ResponseWriter.
	*/
	w.WriteHeader(status)

	/*
		// 5.3 Caching Templates

		// Execute the template set and write the response body. Again, if there
		// is any error we call the the serverError() helper.
	*/
	// err := ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }


	/*
		// 5.4 Catching runtime errors

		// Write the contents of the buffer to the http.ResponseWriter. Note: this
		// is another time where we pass our http.ResponseWriter to a function that
		// takes an io.Writer.
	*/
	buf.WriteTo(w)
}

/*	
	// 5.6 Common dynamic data

	// Create an newTemplateData() helper, which returns a pointer to a templateData
	// struct initialized with the current year. Note that we're not using the
	// *http.Request parameter here at the moment, but we will do later in the book.
*/
func(app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}
