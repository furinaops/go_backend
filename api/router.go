package api

import (
	"net/http"

	"go_backend/api/auth"
	"go_backend/api/home"
	"go_backend/api/status"
	"go_backend/api/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", home.Handler)

	// Public endpoints/routes
	r.HandleFunc("/api/status", status.Handler)
	r.HandleFunc("/api/users/login", auth.LoginHandler).Methods(http.MethodPost)

	// Protected endpoints/routes
	protected := r.NewRoute().Subrouter()
	protected.HandleFunc("/api/users", users.Handler(db))
	protected.Use(auth.AuthMiddleware)

	return r
}
