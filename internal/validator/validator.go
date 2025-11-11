package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)
type Validator struct {
	FieldErrors 	map[string]string

	/*
		// 11.4 User Login

		// Add a new NonFieldErrors []string field to the struct, which we will use to
		// hold any validation errors which are not related to a specific form field.
	*/
	NonFieldErrors	[]string
}

func (v *Validator) Valid() bool {
	// return len(v.FieldErrors) == 0

	/*
		// 11.4 User Login

		// Update the Valid() method to also check that the NonFieldErrors slice is
		// empty
	*/
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {

	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

/*
	// 11.4 User Login

	// Create an AddNonFieldError() helper for adding error messages to the new
	// NonFieldErrors slice.
*/
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

/*
	// 11.3. User signup and password encryption: Validating the user input

	// Use the regexp.MustCompile() function to parse a regular expression pattern
	// for sanity checking the format of an email address. This returns a pointer to
	// a 'compiled' regexp.Regexp type, or panics in the event of an error. 
	// Parsing this pattern once at startup and storing the compiled *regexp.Regexp in a
	// variable is more performant than re-parsing the pattern each time we need it.
*/
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zAZ0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

/*
	// 11.3. User signup and password encryption: Validating the user input

	// MinChars() returns true if a value contains at least n characters.
*/
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

/*
	// 11.3. User signup and password encryption: Validating the user input

	// Matches() returns true if a value matches a provided compiled regular expression pattern.
*/
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}