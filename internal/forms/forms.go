package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// hold all information associated with our form either when it is rendered for the first time,
// or after it is submitted and there might be one or more errors
// creates a custom form struct and embeds a url.Values object
type Form struct {
	url.Values
	Errors errors // errors comes from errors.go
}

// returns true if there are no errors, otherwise return false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// initializes the form struct
func New(data url.Values) *Form {
	return &Form {
		data,
		errors(map[string][]string{}), 
	}
}

// checks for required fields
func (f *Form) Required(fields ...string) { // ...string means you can pass in as many string parameters as you want
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" { // removes any extraneous spaces the user may have filled in by mistake
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	// if any of the specified fields are left blank, display an error on the screen
	if x == "" {
		return false
	}
	return true
}

// checks for string minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field) // gets us the field that we want to check
	if len(x) < length { // checks the length of that string
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

/*
r.PostForm map is populated only for POST, PATCH, and PUT requests, and contains the form data from the request body 
r.Form map is populated for all requests (irrespective of their HTTP method), and contains the form data from any request body and anyh query string 
parameters. In the event on a conflict, the request body value will take precedence over the query string parameters
*/
