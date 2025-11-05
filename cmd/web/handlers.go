package main

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"KanishkaVerma054/snipperBox.dev/internal/validator"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
	// 8.4 Displaying errors and repopulating fields

	// Define a snippetCreateForm struct to represent the form data and
	validation
	// errors for the form fields. Note that all the struct fields are
	deliberately
	// exported (i.e. start with a capital letter). This is because struct fields
	// must be exported in order to be read by the html/template package when
	// rendering the template.
*/
/*
	// 8.6 Automatic form parsing

	// Update our snippetCreateForm struct to include struct tags which tell the
	// decoder how to map HTML form values into the different struct fields. So, for
	// example, here we're telling the decoder to store the value from the HTML form
	// input with the name "title" in the Title field. The struct tag `form:"-"`
	// tells the decoder to completely ignore a field during decoding.
*/
type snippetCreateForm struct {
	Title		string	`form:"title"`
	Content 	string	`form:"content"`
	Expires 	int		`form:"expires"`
	// FieldErrors map[string]string

	/*
		// 8.5 Creating validation helpers

		// Remove the explicit FieldErrors struct field and instead embed the Validator
		// type. Embedding this means that our snippetCreateForm "inherits" all the
		// fields and methods of our Validator type (including the FieldErrors field).
	*/
	validator.Validator		`form:"-"`
}


func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
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

	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.html", data)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display the form for creating a new snippet....."))

	/*
		// 8.1 Setting up a HTML form
	*/
	data := app.newTemplateData(r)

	/*
		// 8.4 Displaying errors and repopulating fields

		// Initialize a new createSnippetForm instance and pass it to the template.
		// Notice how this is also a great opportunity to set any default or
		// 'initial' values for the form --- here we set the initial value for the
	*/
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	/*
		// 8.6 Automatic form parsing

		// Declare a new empty instance of the snippetCreateForm struct.
	*/
	var form snippetCreateForm

	// title := "O snail"
	// content := "O snail\nClimb Mount Fuji, \nBut slowly, slowly!\n\n-Kobayashi Issa"
	// expires := 7

	// id, err := app.snippets.Insert(title, content, expires)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	/*
		// 8.1 Setting up a HTML form

		// First we call r.ParseForm() which adds any data in POST request bodies
		// to the r.PostForm map. This also works in the same way for PUT and PATCH
		// requests. If there are any errors, we use our app.ClientError() helper to
		// send a 400 Bad Request response to the user.
	*/
	// err := r.ParseForm()
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	/*
		// 8.6 Automatic form parsing: Creating a decodePostForm helper

		// Go ahead and update it to use the decodePostForm() helper and remove the r.ParseForm() call, so that the code looks like this:
	*/
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*
		// 8.6 Automatic form parsing

		// Call the Decode() method of the form decoder, passing in the current
		// request and *a pointer* to our snippetCreateForm struct. This will
		// essentially fill our struct with the relevant values from the HTML form.
		// If there is a problem, we return a 400 Bad Request response to the client.
	*/
	// err = app.formDecoder.Decode(&form, r.PostForm)
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	/*
		// 8.1 Setting up a HTML form

		// Use the r.PostForm.Get() method to retrieve the title and content
		// from the r.PostForm map.
	*/
	// title := r.PostForm.Get("title")
	// content := r.PostForm.Get("content")

	/*
		// 8.1 Setting up a HTML form

		// The r.PostForm.Get() method always returns the form data as a *string*.
		// However, we're expecting our expires value to be a number, and want to
		// represent it in our Go code as an integer. So we need to manually covert
		// the form data to an integer using strconv.Atoi(), and we send a 400 Bad
		// Request response if the conversion fails.
	*/
	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	/*
		// 8.2 Setting up a HTML form

		// Initialize a map to hold any validation errors for the form fields.
	*/
	// fieldErrors := make(map[string]string)

	/*
		// 8.3 Validating form data

		// Check that the title value is not blank and is not more than 100
		// characters long. If it fails either of those checks, add a message to the
		// errors map using the field name as the key.
	*/
	// if strings.TrimSpace(title) == "" {
	// 	fieldErrors["title"] = "This field cannot be blank"
	// } else if utf8.RuneCountInString(title) > 100{
	// 	fieldErrors["title"] = "This field cannot be more than 100 characters long"
	// }

	/*
		// 8.3 Validating form data

		// Check that the Content value isn't blank.
	*/
	// if strings.TrimSpace(content) == "" {
	// 	fieldErrors["content"] = "This field cannot be blank"
	// }

	/*
		// 8.3 Validating form data

		// Check the expires value matches one of the permitted values (1, 7 or
		// 365).
	*/
	// if expires != 1 && expires != 7 && expires != 365 {
	// 	fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	// }

	/*
		// 8.3 Validating form data

		// If there are any errors, dump them in a plain text HTTP response and
		// return from the handler.
	*/
	// if len(fieldErrors) > 0 {
	// 	fmt.Fprint(w, fieldErrors)
	// 	return
	// }

	/*
		// 8.4 Displaying errors and repopulating fields

		// Get the expires value from the form as normal.
	*/
	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	// if err != nil {
	// 	app.clientError(w,  http.StatusBadRequest)
	// 	return
	// }

	/*
		// 8.4 Displaying errors and repopulating fields

		// Create an instance of the snippetCreateForm struct containing the values
		// from the form and an empty map for any validation errors.
	*/
	// form := snippetCreateForm{
	// 	Title: r.PostForm.Get("title"),
	// 	Content: r.PostForm.Get("content"),
	// 	Expires: expires,

	// 	/*
	// 		// 8.5 Creating validation helpers

	// 		// Remove the FieldErrors assignment from here.
	// 	*/
	// 	// FieldErrors: map[string]string{},
	// }

	/*
		// 8.4 Displaying errors and repopulating fields

		// Update the validation checks so that they operate on the snippetCreateForm
		// instance.
	*/
	// if strings.TrimSpace(form.Title) == "" {
	// 	form.FieldErrors["title"] = "This field cannot be blank"
	// } else if utf8.RuneCountInString(form.Title) > 100 {
	// 	form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	// }

	// if strings.TrimSpace(form.Content) == "" {
	// 	form.FieldErrors["content"] = "This field cannot be blank"
	// }

	// if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
	// 	form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
	// }

	/*
		// 8.5 Creating validation helpers

		// Because the Validator type is embedded by the snippetCreateForm struct,
		// we can call CheckField() directly on it to execute our validation checks.
		// CheckField() will add the provided key and error message to the
		// FieldErrors map if the check does not evaluate to true. For example, in
		// the first line here we "check that the form.Title field is not blank". In
		// the second, we "check that the form.Title field has a maximum character
		// length of 100" and so on.
	*/
	// form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	// form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	// form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "content", "This field cannot be blank")
	// form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	/*
		// 8.6 Automatic form parsing

		// Then validate and use the data as normal...
	*/
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	/*
		// 8.5 Creating validation helpers

		// Use the Valid() method to see if any of the checks failed. If they did,
		// then re-render the template passing in the form in the same way as
		// before.
	*/
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	/*
		// 8.4 Displaying errors and repopulating fields

		// If there are any validation errors re-display the create.tmpl template,
		// passing in the snippetCreateForm instance as dynamic data in the Form
		// field. Note that we use the HTTP status code 422 Unprocessable Entity
		// when sending the response to indicate that there was a validation error.
	*/
	// if len(form.FieldErrors) > 0 {
	// 	data := app.newTemplateData(r)
	// 	data.Form = form
	// 	app.render(w, http.StatusUnprocessableEntity, "create.html", data)
	// 	return
	// }

	/*
		// 8.4 Displaying errors and repopulating fields

		// If there are any validation errors re-display the create.tmpl template,
		// passing in the snippetCreateForm instance as dynamic data in the Form
		// field. Note that we use the HTTP status code 422 Unprocessable Entity
		// when sending the response to indicate that there was a validation error.
	*/
	// if len(form.FieldErrors) > 0 {
	// 	data := app.newTemplateData(r)
	// 	data.Form = form
	// 	app.render(w, http.StatusUnprocessableEntity, "create.html", data)
	// 	return
	// }

	// id, err := app.snippets.Insert(title, content, expires)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	/*
		// 8.4 Displaying errors and repopulating fields

		// We also need to update this line to pass the data from the
		// snippetCreateForm instance to our Insert() method.
	*/
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
