package forms

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/Somedata", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/SomeData", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form showed valid when required missing found")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "abghany")
	postedData.Add("c", "sss")
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form showed invalid when should be valid")
	}

}
func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	r := httptest.NewRequest("POST", "/SomeData", nil)
	postedData := url.Values{}
	postedData.Add("a", "aboemira@email.com")
	r.PostForm = postedData
	form = New(r.PostForm)
	form.IsEmail("a")
	fmt.Println(form)
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}
}
func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some_value")
	form = New(postedData)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min length of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length if 1 is not met when it is")
	}
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have error but got one")
	}
}
