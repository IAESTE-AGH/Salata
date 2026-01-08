package main

import (
	"go_server/internal/config"
	"go_server/internal/database"
	"go_server/internal/handlers"
	"go_server/middleware"
	"log"
	"net/http"
)

func main() {
	database.Init(config.DatabaseURL)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /status", handlers.HandleGetStatus)
	mux.HandleFunc("GET /users", handlers.HandleGetUsers)
	mux.HandleFunc("POST /create_user", handlers.HandleCreateUser)
	mux.HandleFunc("GET /verify_user", handlers.HandleVerifyUser)
	mux.HandleFunc("POST /login", handlers.HandleLogin)

	privateMux := http.NewServeMux()

	privateMux.HandleFunc("POST /create_event", handlers.HandleCreateEvent)
	privateMux.HandleFunc("POST /delete_event", handlers.HandleDeleteEvent)
	privateMux.HandleFunc("POST /change_availability", handlers.HandleChangeAvailability)
	privateMux.HandleFunc("POST /get_all_current_events", handlers.HandleGetAllCurrentEvents)

	mux.Handle("/", middleware.AuthMiddleware(privateMux))

	log.Println("Server is running at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
