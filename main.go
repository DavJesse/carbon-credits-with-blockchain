package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	// Registration logic
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	// Login logic
}

func main() {
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", loginUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
