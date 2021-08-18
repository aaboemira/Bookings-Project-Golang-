package forms

import "fmt"

type errors map[string][]string

// adds an error messages for a given field
func (e errors) add(field, message string) {
	e[field] = append(e[field], message)
}

// returns first error message
func (e errors) Get(field string) string {
	fmt.Println(field)
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
