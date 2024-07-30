package main

import (
	"carbo-cred/backend/database"
	"carbo-cred/backend/user"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	http.Handle("frontend/site/assets/css", http.StripPrefix("frontend/site/assets/css", http.FileServer(http.Dir("/home/johnotieno0/hackathon/carbon-credits-with-blockchain/frontend/site/assets/css"))))
	http.HandleFunc("/", user.HomeHandler)
	http.HandleFunc("/register", user.RegisterUser)
	http.HandleFunc("/login", user.LoginUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
