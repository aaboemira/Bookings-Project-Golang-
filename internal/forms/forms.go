package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	Errors errors
	url.Values
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		errors(map[string][]string{}),
		data,
	}
}
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.add(field, "This field cannot be blank")
			fmt.Println(f.Errors)
		}
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := f.Get(field)
	if x == "" {
		f.Errors.add(field, "This field can't be empty")
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.add(field, "Invalid email address")
	}
}

func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}
