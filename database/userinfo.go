package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// var db *sql.DB

// func initDB() {
//     var err error
//     connStr := "user=username dbname=yourdbname password=yourpassword sslmode=disable"
//     db, err = sql.Open("postgres", connStr)
//     if err != nil {
//         log.Fatal(err)
//     }

//     _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
//         id SERIAL PRIMARY KEY,
//         username TEXT UNIQUE NOT NULL,
//         password TEXT NOT NULL
//     )`)
//     if err != nil {
//         log.Fatal(err)
//     }
// }
