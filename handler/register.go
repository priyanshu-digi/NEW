package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// User represents a user in the database
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler handles POST requests to the "/register" endpoint
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body to get the user's registration information
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Insert the user's information into the database
		query := `
			INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)
			RETURNING id
		`
		var id int
		if err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set the response status to 201 Created and write the user ID to the response body
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}
