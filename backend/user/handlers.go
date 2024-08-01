package user

import (
	"context"
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"carbo-cred/backend/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)



// Templates
var (
    registerTmpl = template.Must(template.ParseFiles("login/login.html"))
    loginTmpl = template.Must(template.ParseFiles("login/login.html"))
    homeTmpl = template.Must(template.ParseFiles("index.html"))
    successTmpl = template.Must(template.ParseFiles("frontend/success/success.html"))
    errorTmpl = template.Must(template.ParseFiles("frontend/error/error.html"))

)

func ServeHomePage(w http.ResponseWriter, r *http.Request) {
    homeTmpl.Execute(w, nil)
}

// RegisterUser handles user registration and shows registration page
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		registerTmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.Collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	successTmpl.Execute(w, nil)  // Display success page
}

// LoginUser handles user login and shows login page
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		loginTmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": user.Username}
	err := database.Collection.FindOne(ctx, filter).Decode(&storedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusUnauthorized)
			errorTmpl.Execute(w, nil)  // Display error page
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		errorTmpl.Execute(w, nil)  // Display error page
		return
	}

	w.WriteHeader(http.StatusOK)
	successTmpl.Execute(w, nil)  // Display success page
}
