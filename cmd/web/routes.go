package main

import "net/http"

/*
	// 3.5: Isolating the application routes

	// The routes() method returns a servemux containing our application routes.
*/
func(app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	filerServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}