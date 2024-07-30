package user

import (
	"context"
	"encoding/json"
<<<<<<< HEAD
	"html/template"
	"net/http"
=======
	"net/http"
	"text/template"
>>>>>>> master
	"time"

	"carbo-cred/backend/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

<<<<<<< HEAD
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("/home/johnotieno0/hackathon/carbon-credits-with-blockchain/frontend/site/index.html"))
    tmpl.Execute(w, nil)
}

// RegisterUser handles user registration
=======
// Get the absolute path of the templates
// func getTemplatePath(relativePath string) string {
// 	_, b, _, _ := runtime.Caller(0)
// 	basePath := filepath.Dir(b)
// 	return filepath.Join(basePath, relativePath)
// }

// Templates
var registerTmpl = template.Must(template.ParseFiles("/home/davodhiambo/github/carbon-credits-with-blockchain/server/register.html"))
var loginTmpl = template.Must(template.ParseFiles("/home/davodhiambo/github/carbon-credits-with-blockchain/login/login.html"))

// User struct
// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// RegisterUser handles user registration and shows registration page
>>>>>>> master
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
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
