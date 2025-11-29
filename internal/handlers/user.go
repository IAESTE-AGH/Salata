package handlers

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/repository"
	"net/http"
)

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	users, _ := repository.GetAllUsers(db)
	json.NewEncoder(w).Encode(users)
}
