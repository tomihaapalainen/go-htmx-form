package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

type PageData struct {
	FirstName      string
	LastName       string
	Email          string
	FirstNameError string
	LastNameError  string
	EmailError     string
	PasswordError  string
}

func HandleIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet:
			handleGetIndex(w, r)
		case r.Method == http.MethodPost:
			handlePostIndex(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func handleGetIndex(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/form.html",
		"templates/form2.html",
		"templates/form-email.html",
		"templates/form-first-name.html",
		"templates/form-last-name.html",
		"templates/form-password.html",
	))
	if err := t.ExecuteTemplate(w, "index", nil); err != nil {
		log.Println("err executing index.html: ", err)
	}
}

func handlePostIndex(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hasError := false

	pageData := PageData{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	firstName = strings.TrimSpace(firstName)
	if firstName == "" {
		pageData.FirstNameError = "first name must not be empty"
		hasError = true
	}
	lastName = strings.TrimSpace(lastName)
	if lastName == "" {
		pageData.LastNameError = "last name must not be empty"
		hasError = true
	}
	email = strings.TrimSpace(email)
	if email == "" {
		pageData.EmailError = "email must not be empty"
		hasError = true
	}
	re := regexp.MustCompile(`^.+@.+\..+$`)
	if !re.Match([]byte(email)) {
		pageData.EmailError = "email must be a valid email address"
		hasError = true
	}
	if !passwordIsValid(password) {
		pageData.PasswordError = `password must have at least 8 characters, and contain
			1 uppercase and 1 lowercase character, and a number`
		hasError = true
	}

	if hasError {
		t := template.Must(template.ParseFiles(
			"templates/form.html",
		))

		if err := t.ExecuteTemplate(w, "form", pageData); err != nil {
			log.Println("err executing form.html: ", err)
		}
	}
	//...
}

func HandlePostFirstName(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")

	firstName = strings.TrimSpace(firstName)

	pageData := PageData{
		FirstName: firstName,
	}
	if firstName == "" {
		pageData.FirstNameError = "first name must not be empty"
	}

	t := template.Must(template.ParseFiles(
		"templates/form-first-name.html",
	))
	if err := t.ExecuteTemplate(w, "form-first-name", pageData); err != nil {
		log.Println("err executing form-first-name.html: ", err)
	}
}

func HandlePostLastName(w http.ResponseWriter, r *http.Request) {
	lastName := r.FormValue("last_name")

	lastName = strings.TrimSpace(lastName)

	pageData := PageData{
		LastName: lastName,
	}
	if lastName == "" {
		pageData.LastNameError = "enter a valid email address"
	}

	t := template.Must(template.ParseFiles(
		"templates/form-last-name.html",
	))
	if err := t.ExecuteTemplate(w, "form-last-name", pageData); err != nil {
		log.Println("err executing form-last-name.html: ", err)
	}
}

func HandlePostEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	email = strings.TrimSpace(email)

	pageData := PageData{
		Email: email,
	}
	re := regexp.MustCompile(`^.+@.+\..+$`)
	if !re.Match([]byte(email)) {
		pageData.EmailError = "email must be a valid email address"
	}

	t := template.Must(template.ParseFiles(
		"templates/form-email.html",
	))
	if err := t.ExecuteTemplate(w, "form-email", pageData); err != nil {
		log.Println("err executing form-email.html: ", err)
	}
}

func HandlePostPassword(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")

	password = strings.TrimSpace(password)

	pageData := PageData{
		FirstName: password,
	}
	if !passwordIsValid(password) {
		pageData.PasswordError = `password must have at least 8 characters, and contain
		1 uppercase and 1 lowercase character, and a number`
	}

	t := template.Must(template.ParseFiles(
		"templates/form-password.html",
	))
	if err := t.ExecuteTemplate(w, "form-password", pageData); err != nil {
		log.Println("err executing form-password.html: ", err)
	}
}

func passwordIsValid(password string) bool {
	count := 0
	hasUppercase := false
	hasLowercase := false
	hasNumber := false

	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			hasUppercase = true
		}

		if r >= 'a' && r <= 'z' {
			hasLowercase = true
		}

		if r >= '0' && r <= '9' {
			hasNumber = true
		}

		count++
	}

	return count >= 8 && hasUppercase && hasLowercase && hasNumber
}
