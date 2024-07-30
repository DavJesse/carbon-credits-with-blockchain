package main

import (
	"carbo-cred/db"
	"carbo-cred/user"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.LoginUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
