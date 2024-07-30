package main

import (
	"carbo-cred/backend/database"
	"carbo-cred/backend/user"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	http.HandleFunc("/register", user.ShowRegisterPage)
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.ShowLoginPage)
	http.HandleFunc("/login", user.LoginUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
