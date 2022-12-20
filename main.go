package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/priyanshu-digi/NEW/handler"
)

func main() {
	// Connect to the database
	connStr := "postgresql://username:password@localhost/database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set up the HTTP server
	http.HandleFunc("/register", handler.RegisterHandler(db))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
