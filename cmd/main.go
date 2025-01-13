package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/refalah/go-dashboard-be/internal/user"
	"github.com/refalah/go-dashboard-be/web"
)

func main() {
	// Initialize database connection
	db, err := web.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer db.Close()

	// Initialize router
	r := web.NewRouter()

	// Register user routes
	user.RegisterRoutes(r, db)

	// Start the web server
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
