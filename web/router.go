package web

import (
	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

// NewRouter initializes the router for the application
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

// InitDB sets up the PostgreSQL connection
func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=admin123 dbname=unodb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
