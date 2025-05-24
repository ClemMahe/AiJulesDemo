package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// User struct to hold username and password hash
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"-"` // Do not expose hash in responses
}

// Credentials struct for request payloads
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	users      = make(map[string]User)
	usersMutex = &sync.Mutex{} // To protect concurrent access to the users map
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	passwordHash, err := hashPassword(creds.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	users[creds.Username] = User{Username: creds.Username, PasswordHash: passwordHash}
	log.Printf("User %s registered successfully", creds.Username)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	user, exists := users[creds.Username]
	usersMutex.Unlock()

	if !exists {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(creds.Password, user.PasswordHash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in successfully", creds.Username)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func main() {
	http.HandleFunc("/auth/register", registerHandler)
	http.HandleFunc("/auth/login", loginHandler)

	port := "8080"
	log.Printf("Authentication service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
