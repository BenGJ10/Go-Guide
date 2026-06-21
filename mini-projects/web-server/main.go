package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var userCache = make(map[int]User)

var cacheMutex sync.RWMutex

func main() {
	// Create a multiplexer to handle incoming HTTP requests
	mux := http.NewServeMux()

	// Define a handler function for the root path
	mux.HandleFunc("/", handleRoot)

	// Define a handler function for the "/users" path to handle POST requests for creating users
	mux.HandleFunc("POST /users", createUser)

	// Define a handler function for the "/users" path to handle GET requests for retrieving all users
	mux.HandleFunc("GET /users", getAllUsers)

	// Define a handler function for the "/users/{id}" path to handle GET requests for retrieving user information
	mux.HandleFunc("GET /users/{id}", getUser)

	// Define a handler function for the "/users/{id}" path to handle DELETE requests for deleting users
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server listening to :8080")
	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", mux)
}

// handleRoot is the handler function for the root path ("/"). It writes a welcome message to the response.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Ben.. Welcome to the server!")
}

// createUser is the handler function for the "/users" path. It handles POST requests to create a new user.
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// Decode the JSON request body into a User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age <= 0 || user.Age > 120 {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// Store the user in the cache (using a simple ID based on the current size of the cache)
	// Use a mutex to ensure thread-safe access to the userCache
	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	// Respond with a 204 No Content status to indicate successful creation of the user
	w.WriteHeader(http.StatusNoContent)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Use a read lock to safely access the userCache while encoding it to JSON
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the userCache as JSON and write it to the response
	err := json.NewEncoder(w).Encode(userCache)
	if err != nil {
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
		return
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Set the Content-Type header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the user data as JSON and write it to the response
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	_, ok := userCache[id]
	if !ok {
		cacheMutex.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
