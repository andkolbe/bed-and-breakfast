package forms

type errors map[string][]string // values in the map are a slice of strings

// adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message) // append errors to the slice for a given field with a message
}

// returns the first error message
func (e errors) Get(field string) string {
	es := e[field] // assign es the values of whatever errors we find in our map
	if len(es) == 0 { // if we have no errors
		return ""
	}
	return es[0] // if there are multiple errors, display the first one
}