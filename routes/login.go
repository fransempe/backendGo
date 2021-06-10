package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fransempe/backendGo/jwt"
	"github.com/fransempe/backendGo/models"
)

// Login is a function that is responsible for making the user login
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "asdadsdsa")
	w.Header().Add("content-type", "application/json")

	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid User or password "+err.Error(), 400)
		return
	}

	//dbDoc, found := db.TryLogin(newUser.Username, newUser.Password)
	var result models.User
	var found = false

	if newUser.Username == "admin" && newUser.Password == "12345" {
		found = true
	}

	if !found {
		http.Error(w, "Invalid user or password ", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(result)

	if err != nil {
		http.Error(w, "Failed in generating token "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

	//save token in a cookie
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
