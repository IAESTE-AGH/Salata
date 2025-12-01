package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

type RequestDeleteEvent struct {
	EventId int `json:"event_id"`
}

func HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req RequestDeleteEvent
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.DeleteEvent(db, req.EventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
