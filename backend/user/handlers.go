package user

import (
	"net/http"
	"text/template"
)

var registerTmpl = template.Must(template.ParseFiles("../frontend/register.html"))
var loginTmpl = template.Must(template.ParseFiles("../frontend/login.html"))

// ShowRegisterPage renders the registration page
func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	registerTmpl.Execute(w, nil)
}

// ShowLoginPage renders the login page
func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	loginTmpl.Execute(w, nil)
}
