package main

import (
	"fmt"
	"log"
	"net/http"

	"go_backend/api"
	"go_backend/internal/database"
	"go_backend/test"
)

func main() {
	connStr := "postgres://username:password@localhost:5432/database?sslmode=disable"

	db, err := database.Connect(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handler := api.NewRouter(db)

	port := ":8080"
	fmt.Println("server is up on http://localhost:8080")

	err = http.ListenAndServe(port, test.Logger(handler))
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
