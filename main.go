package main

import (
	"carbo-cred/server"
	"fmt"
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
	//fmt.Fprintf(w, "log in", http.StatusAccepted)
	// Registration logic
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	// Login logic
}

func main() {
	server.InitDB()

	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", loginUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
