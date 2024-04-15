package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var registeredUsers []User

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8081", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registeredUsers = append(registeredUsers, user)
	fmt.Fprintf(w, "User %s registered successfully!", user.Email)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, registeredUser := range registeredUsers {
		if registeredUser.Email == user.Email && registeredUser.Password == user.Password {
			fmt.Fprintf(w, "OK")
			return
		}
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
