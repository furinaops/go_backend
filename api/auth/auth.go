package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go_backend/internal/models"
)

type contextKey string

const userContextKey contextKey = "user"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// TODO: Check if user exists and password is correct
	if user.Name == "admin" && user.Role == "admin" {
		token, err := GenerateToken(user)
		if err != nil {
			http.Error(w, "Failed to generate Token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
		return
	}
	http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// remove bearer prefix (if present)
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		// Add user info to context for later use
		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
