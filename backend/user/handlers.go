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

var registerTmpl = template.Must(template.ParseFiles("../frontend/register.html"))
var loginTmpl = template.Must(template.ParseFiles("../frontend/login.html"))
