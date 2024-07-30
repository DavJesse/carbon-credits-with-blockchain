package user

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"carbo-cred/backend/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("/home/johnotieno0/hackathon/carbon-credits-with-blockchain/frontend/site/index.html"))
    tmpl.Execute(w, nil)
}

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
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

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
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
