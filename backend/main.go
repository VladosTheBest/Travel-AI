package main

import (
	structures "backend/structures"
	"encoding/json"
	"fmt"
	"net/http"
)

var registeredUsers []structures.User

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		if r.Method == "OPTIONS" {
			return
		}
	})

	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	var user structures.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registeredUsers = append(registeredUsers, user)
	fmt.Fprintf(w, "User %s registered successfully!", user.Email)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user structures.User
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

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if req.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
