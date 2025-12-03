package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
	"time"
)

type RequestCreateEvent struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Name  string    `json:"event_name"`
}

func HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var req RequestCreateEvent
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.CreateEvent(db, req.Start, req.End, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
