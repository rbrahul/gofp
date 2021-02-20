package gofp

import (
	"strings"
	"testing"
)

var mockedUser = map[string]interface{}{
	"name": "John Doe",
	"age":  30,
	"contacts": map[string]interface{}{
		"email":  "johndoe@gmail.com",
		"office": "Google Inc.",
		"fax": map[string]interface{}{
			"uk": "+44-208-1234567",
		},
	},
}

func Test_Pipe(t *testing.T) {
	mockedUser := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"contacts": map[string]interface{}{
			"email":  "johndoe@gmail.com",
			"office": "Google Inc.",
			"fax": map[string]interface{}{
				"uk": "+44-208-1234567",
			},
		},
	}
	getContacts := func(data interface{}) interface{} {
		return data.(map[string]interface{})["contacts"]
	}

	getEmail := func(data interface{}) interface{} {
		return data.(map[string]interface{})["email"]
	}
	getUpperCaseEmail := func(data interface{}) interface{} {
		return strings.ToUpper(data.(string))
	}

	email := Pipe(
		getContacts,
		getEmail,
		getUpperCaseEmail,
	)(mockedUser)
	if email != "JOHNDOE@GMAIL.COM" {
		t.Errorf("Pipe() = %v, want %v", email, "JOHNDOE@GMAIL.COM")
	}
}

func Test_Compose(t *testing.T) {
	getContacts := func(data interface{}) interface{} {
		return data.(map[string]interface{})["contacts"]
	}

	getEmail := func(data interface{}) interface{} {
		return data.(map[string]interface{})["email"]
	}
	getUpperCaseEmail := func(data interface{}) interface{} {
		return strings.ToUpper(data.(string))
	}

	email := Compose(
		getUpperCaseEmail,
		getEmail,
		getContacts,
	)(mockedUser)
	if email != "JOHNDOE@GMAIL.COM" {
		t.Errorf("Pipe() = %v, want %v", email, "JOHNDOE@GMAIL.COM")
	}
}
