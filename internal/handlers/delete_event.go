package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

type Request struct {
	EventId int `json:"event_id"`
}

func HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req Request
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
