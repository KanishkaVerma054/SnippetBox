package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	// "log"
)

/*
	3.3 Dependency Injection
	// Change the signature of the home handler so it is defined as a method against
	// *application.
*/

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w) // 3.4 Centralized Error: Use the notFound() helper
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Print(err.Error()) // This isn't using our new error logger.

	/*
	// 3.3 Dependency Injection 

		// Because the home handler function is now a method against application
		// it can access its fields, including the error logger. We'll write the log
		// message to this instead of the standard logger.
	*/
		// app.errorLog.Print(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		/*
			// 3.4 Centralized Error:
		*/
		app.serverError(w, err) //Use the serverError() helper.

		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.Print(err.Error())

		/*
			3.3 Dependency Injection

			// Also update the code here to use the error logger from the application
			// struct.
		*/
		// app.errorLog.Print(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		/*
			// 3.4 Centralized Error:
		*/
		app.serverError(w, err) // Use the serverError() helper.


	}
}

/*
// 3.3 Dependency Injection

	// Change the signature of the snippetView handler so it is defined as a method
	// against *application.
*/

func(app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)

		/*
			3.4 Centalized Error:
		*/
		app.notFound(w) // Use the notFound() helper.
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

/*
// 3.3 Dependency Injection

	// Change the signature of the snippetView handler so it is defined as a method
	// against *application.
*/

func(app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		/*
			3.4 Centalized Error:
		*/
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		
		return
	}

	w.Write([]byte("Creating a new snippet"))
}