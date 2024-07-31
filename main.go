package main

import (
	"carbo-cred/backend/database"
	"carbo-cred/backend/user"
	"net/http"
)

func main() {
	database.InitDB()
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("login"))))
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/", user.ServeHomePage)
	http.HandleFunc("/login", user.LoginUser)

	http.ListenAndServe(":8080", nil)
}
