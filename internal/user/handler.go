package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes will register user routes
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/users", GetUsers(db)).Methods("GET")
	// r.HandleFunc("/users/{id}", GetUser(db)).Methods("GET")
	r.HandleFunc("/users", CreateUser(db)).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser(db)).Methods("DELETE")
}

// GetUsers retrieves all users from the database
func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, first_name, last_name, email FROM users_test")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var firstName, lastName, email string
			if err := rows.Scan(&id, &firstName, &lastName, &email); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, map[string]interface{}{
				// "id":         id,
				"first_name": firstName,
				"last_name":  lastName,
				"email":      email,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// GetUser retrieves a single user from the database
// func GetUser(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := mux.Vars(r)["id"]
// 		row := db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id=$1", id)

// 		var user map[string]interface{}
// 		if err := row.Scan(&user["id"], &user["first_name"], &user["last_name"], &user["email"]); err != nil {
// 			http.Error(w, "User not found", http.StatusNotFound)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(user)
// 	}
// }

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO users_test(first_name, last_name, email) VALUES($1, $2, $3)", user["first_name"], user["last_name"], user["email"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User created")
	}
}

// UpdateUser updates an existing user in the database
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var user map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4", user["first_name"], user["last_name"], user["email"], id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User updated")
	}
}

// DeleteUser deletes a user from the database
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User deleted")
	}
}
