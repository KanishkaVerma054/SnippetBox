package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form"
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

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)

	buf.WriteTo(w)
}

func(app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

/*
	// 8.6 Automatic form parsing: Creating a decodePostForm helper

	// Create a new decodePostForm() helper method. The second parameter here, dst,
	// is the target destination that we want to decode the form data into.
*/
func(app *application) decodePostForm(r *http.Request, dst any) error {
	/*
		// 8.6 Automatic form parsing: Creating a decodePostForm helper

		// Call ParseForm() on the request, in the same way that we did in our
		// createSnippetPost handler.
	*/
	err := r.ParseForm()
	if err != nil {
		return err
	}

	/*
		// 8.6 Automatic form parsing: Creating a decodePostForm helper

		// Call Decode() on our decoder instance, passing the target destination as
		// the first parameter.
	*/
	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		/*
			// 8.6 Automatic form parsing: Creating a decodePostForm helper

			// If we try to use an invalid target destination, the Decode() method
			// will return an error with the type *form.InvalidDecoderError.We use
			// errors.As() to check for this and raise a panic rather than returning
			// the error.
		*/
		var InvalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &InvalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}