package main

import (
	"net/http"

	"github.com/justinas/alice"
)

/*
	// 6.2 Setting security headers

	// Update the signature for the routes() method so that it returns a
	// http.Handler instead of *http.ServeMux.
*/
func(app *application) routes() http.Handler {
	mux := http.NewServeMux()

	filerServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// return mux

	/*
		// 6.2 Setting security headers

		// Pass the servemux as the 'next' parameter to the secureHeaders middleware.
		// Because secureHeaders is just a function, and the function returns a
		// http.Handler we don't need to do anything else.
	*/
	// return secureHeaders(mux)

	// /*
	// 	// 6.3 Request Logging

	// 	// Wrap the existing chain with the logRequest middleware.
	// */
	// return app.logRequest(secureHeaders(mux))

	/*
		// 6.4 Panic Recovery

		// Wrap the existing chain with the recoverPanic middleware.
	*/
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	/*
		// 6.5 Composable middleware chains

		// Create a middleware chain containing our 'standard' middleware
		// which will be used for every request our application receives.
	*/
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	/*
		// 6.5 Composable middleware chains

		// Return the 'standard' middleware chain followed by the servemux.
	*/
	return standard.Then(mux)
}