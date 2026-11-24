package users

import (
	"encoding/json"
	"log"
	"net/http"

	"go_backend/internal/models"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			http.Error(w, "Failed to query users", http.StatusInternalServerError)
			log.Printf("database query error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
