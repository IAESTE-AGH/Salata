package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"go_server/middleware"
	"net/http"
)

type RequestChangeAvailability struct {
	EventId      int               `json:"event_id"`
	Availability map[string]string `json:"availability"`
}

func HandleChangeAvailability(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	var req RequestChangeAvailability
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.DB

	err = repository.ChangeAvailability(db, req.EventId, userID, req.Availability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Availability updated"))
}
