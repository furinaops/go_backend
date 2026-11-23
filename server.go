package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// NOTE: not implemented for now — changed my mind, using GORM instead.
	// "database/sql"
	// _ "github.com/jackc/pgx/v5/stdlib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// NOTE: not implemented for now — changed my mind, using GORM instead.
// var db *sql.DB
var db *gorm.DB

func main() {

	// NOTE: not implemented for now — changed my mind, using GORM instead.
	// var err error
	// connStr := "postgres://username:password@localhost:5432/database?sslmode=disable"
	//
	// db, err = sql.Open("pgx", connStr)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database: ", err)
	// }
	// defer db.Close()
	//
	// if err = db.Ping(); err != nil {
	// 	// FIX: Fatalf had no format verb, so the error was silently dropped. Added %v.
	// 	log.Fatalf("Failed to ping database: %v", err)
	// }

	var err error
	connStr := "postgres://username:password@localhost:5432/database?sslmode=disable"

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to the server!")
	})

	// NOTE: not implemented for now — changed my mind, using GORM instead.
	// mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
	//
	// 	rows, err := db.Query("SELECT id, name, role FROM users")
	// 	if err != nil {
	// 		// FIX: http.Error requires a status code; only 2 args were passed. Added 500.
	// 		http.Error(w, "Failed to query users", http.StatusInternalServerError)
	// 		// FIX: %v verb was used but err was not passed, so the error was never logged.
	// 		log.Printf("database query error: %v", err)
	// 		return
	// 	}
	// 	defer rows.Close()
	//
	// 	var users []User
	//
	// 	for rows.Next() {
	// 		var u User
	// 		if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
	// 			// FIX: missing status code argument (and space in `err:=`); gofmt'd.
	// 			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
	// 			log.Printf("database scan error: %v", err)
	// 			return
	// 		}
	//
	// 		users = append(users, u)
	// 	}
	//
	// 	if err = rows.Err(); err != nil {
	// 		// FIX: missing status code argument; block also had broken indentation.
	// 		http.Error(w, "Error iteration over rows", http.StatusInternalServerError)
	// 		// FIX: %v verb was used but err was not passed.
	// 		log.Printf("Row iteration error: %v", err)
	// 		return
	// 	}
	//
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(users)
	// })

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		var users []User

		if err := db.Find(&users).Error; err != nil {
			http.Error(w, "Failed to query users", http.StatusInternalServerError)
			log.Printf("database query error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		status := map[string]string{
			"status":  "online",
			"version": "1.0.0",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	})

	wrappedMux := Logger(mux)

	port := ":8080"
	fmt.Println("server is up on http://localhost:8080")

	// FIX: used := but err is already declared above ("no new variables on left side of :=").
	err = http.ListenAndServe(port, wrappedMux)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
