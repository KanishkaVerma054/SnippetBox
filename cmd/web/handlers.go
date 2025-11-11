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

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

/*
// 11.3. User signup and password encryption

// Create a new userSignupForm struct.
*/
type userSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"_"`
}

/*
	// 11.4 User Login

	// Create a new userLoginForm struct.
*/
type userLoginForm struct {
	Email				string `form:"email"`
	Password			string `form:"password"`
	validator.Validator	`form:"_"`
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
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

/*
// 11.1 Routes setup
*/
func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	/*
		// 11.3. User signup and password encryption

		// Update the handler so it displays the signup page.
	*/
	data := app.newTemplateData(r)
	data.Form = userSignupForm{}
	app.render(w, http.StatusOK, "signup.html", data)
}
func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	/*
		// 11.3. User signup and password encryption: Validating the user input

		// Declare an zero-valued instance of our userSignupForm struct.
	*/
	var form userSignupForm

	/*
		// 11.3. User signup and password encryption: Validating the user input

		// Parse the form data into the userSignupForm struct.
	*/
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*
		// 11.3. User signup and password encryption: Validating the user input

		// Validate the form contents using our helper functions.
	*/
	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

	/*
		// 11.3. User signup and password encryption: Validating the user input

		// If there are any errors, redisplay the signup form along with a 422 status code.
	*/
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
		return
	}

	/*
		// 11.3. User signup and password encryption: Validating the user input

		// Otherwise send the placeholder response (for now!).
	*/
	// fmt.Fprintln(w, "Create a new user...")

	/*
		// User signup and password encryption: Storing the user details

		// Try to create a new user record in the database. If the email already
		// exists then add an error message to the form and re-display it.
	*/
	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "signup.tmpl",
				data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	/*
		// User signup and password encryption: Storing the user details

		// Otherwise add a confirmation flash message to the session confirming that
		// their signup worked.
	*/
	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	/*
		// User signup and password encryption: Storing the user details

		// And redirect the user to the login page.
	*/
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display a HTML form for logging in a user...")

	/*
		// 11.4 User Login

		// Update the handler so it displays the login page.
	*/
	data := app.newTemplateData(r)
	data.Form = userLoginForm{}
	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Authenticate and login the user...")

	/*
		// 11.4 User Login: Verifying the user details

		// Decode the form data into the userLoginForm struct.
	*/
	var form userLoginForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*
		// 11.4 User Login: Verifying the user details

		// Do some validation checks on the form. We check that both email and
		// password are provided, and also check the format of the email address as
		// a UX-nicety (in case the user makes a typo).
	*/
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		return
	}

	/*
		// 11.4 User Login: Verifying the user details

		// Check whether the credentials are valid. If they're not, add a generic
		// non-field error message and re-display the login page.
	*/
	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email or password is incorrect")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	/*
		// 11.4 User Login: Verifying the user details

		// Use the RenewToken() method on the current session to change the session
		// ID. It's good practice to generate a new session ID when the
		// authentication state or privilege levels changes for the user (e.g. login
		// and logout operations).
	*/
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	/*
		// 11.4 User Login: Verifying the user details

		// Redirect the user to the create snippet page.
	*/
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Logout the user...")

	/*
		// 11.5 User logout

		// Use the RenewToken() method on the current session to change the session
		// ID again.
	*/
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	/*
		// 11.5 User logout

		// Remove the authenticatedUserID from the session data so that the user is
		// 'logged out'.
	*/
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	/*
		// 11.5 User logout

		// Add a flash message to the session to confirm to the user that they've been
		// logged out.
	*/
	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	/*
		// 11.5 User logout

		// Redirect the user to the application home page.
	*/
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
