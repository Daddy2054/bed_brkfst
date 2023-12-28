package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}




func TestNew(t *testing.T) {
	data := url.Values{}
	data.Add("email", "test@example.com")
	data.Add("password", "secret")

	form := New(data)

	if form.Values["email"][0] != "test@example.com" {
		t.Errorf("form.Values[\"email\"][0] got %q, want %q", form.Values["email"][0], "test@example.com")
	}

	if form.Values["password"][0] != "secret" {
		t.Errorf("form.Values[\"password\"][0] got %q, want %q", form.Values["password"][0], "secret")
	}
}

func TestRequired(t *testing.T) {
	data := url.Values{}
	data.Add("email", "")
	data.Add("password", "")

	form := New(data)
	form.Required("email", "password")

	if len(form.Errors) != 2 {
		t.Errorf("len(form.Errors) got %d, want %d", len(form.Errors), 2)
	}

	if form.Errors["email"][0] != "This field cannot be blank" {
		t.Errorf("form.Errors[\"email\"][0] got %q, want %q", form.Errors["email"][0], "This field cannot be blank")
	}

	if form.Errors["password"][0] != "This field cannot be blank" {
		t.Errorf("form.Errors[\"password\"][0] got %q, want %q", form.Errors["password"][0], "This field cannot be blank")
	}
}

// func TestHas(t *testing.T) {
// 	data := url.Values{}
// 	data.Add("email", "test@example.com")
// 	data.Add("password", "")

// 	form := New(data)

// 	if !form.Has("email", &http.Request{}) {
// 		t.Errorf("form.Has(\"email\", &http.Request{}) got %t, want %t", form.Has("email", &http.Request{}), true)
// 	}

// 	if form.Has("password", &http.Request{}) {
// 		t.Errorf("form.Has(\"password\", &http.Request{}) got %t, want %t", form.Has("password", &http.Request{}), false)
// 	}
// }

// func TestMinLength(t *testing.T) {
// 	data := url.Values{}
// 	data.Add("password", "secret")

// 	form := New(data)
// 	form.MinLength("password", 8, &http.Request{})

// 	if len(form.Errors) != 0 {
// 		t.Errorf("len(form.Errors) got %d, want %d", len(form.Errors), 0)
// 	}

// 	data = url.Values{}
// 	data.Add("password", "s")

// 	form = New(data)
// 	form.MinLength("password", 8, &http.Request{})

// 	if len(form.Errors) != 1 {
// 		t.Errorf("len(form.Errors) got %d, want %d", len(form.Errors), 1)
// 	}

// 	if form.Errors["password"][0] != "This field must be at least 8 characters long" {
// 		t.Errorf("form.Errors[\"password\"][0] got %q, want %q", form.Errors["password"][0], "This field must be at least 8 characters long")
// 	}
// }

func TestIsEmail(t *testing.T) {
	data := url.Values{}
	data.Add("email", "test@example.com")

	form := New(data)
	form.IsEmail("email")

	if len(form.Errors) != 0 {
		t.Errorf("len(form.Errors) got %d, want %d", len(form.Errors), 0)
	}

	data = url.Values{}
	data.Add("email", "invalid")

	form = New(data)
	form.IsEmail("email")

	if len(form.Errors) != 1 {
		t.Errorf("len(form.Errors) got %d, want %d", len(form.Errors), 1)
	}

	if form.Errors["email"][0] != "Invalid email address" {
		t.Errorf("form.Errors[\"email\"][0] got %q, want %q", form.Errors["email"][0], "Invalid email address")
	}
}
