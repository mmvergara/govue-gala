package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mmvergara/golang-gala/backend/model"
	"github.com/mmvergara/golang-gala/backend/repository/user"
	"github.com/mmvergara/golang-gala/backend/utilities"
)



type User struct {
	Repo *user.RedisRepo
}


func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	
	// Create a struct to hold the request body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode the request body into the struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Error decoding body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utilities.HashPassword(body.Password)
	if err != nil {
		fmt.Println("Error hashing password", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a new user
	user := model.User{
		UserID:   uuid.Must(uuid.NewRandom()),
		Username: body.Username,
		Password: hashedPassword,
	}

	// Insert the user into the database
	if err := u.Repo.Insert(r.Context(), user); err != nil {
		fmt.Println("Error inserting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	// Marshal the user into JSON
	res, err := json.Marshal(user.UserID)
	if err != nil {
		fmt.Println("Error marshalling user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
