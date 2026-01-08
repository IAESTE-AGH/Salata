package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"go_server/middleware"
	"net/http"
)

func HandleGetAllCurrentEvents(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	db := database.DB

	eventsList, err := repository.GetAllCurrentEvents(db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(eventsList)
}
