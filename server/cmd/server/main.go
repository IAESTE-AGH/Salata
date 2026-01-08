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

	protect := func(handlerFunc http.HandlerFunc) http.Handler {
		handler := http.HandlerFunc(handlerFunc)
		return middleware.AuthMiddleware(handler)
	}

	// Public (does not require authentication)
	mux.HandleFunc("/status", handlers.HandleGetStatus)
	mux.HandleFunc("/users", handlers.HandleGetUsers)
	mux.HandleFunc("/login", handlers.HandleLogin)
	mux.HandleFunc("/verify_user", handlers.HandleVerifyUser)
	mux.HandleFunc("/create_user", handlers.HandleCreateUser)

	// Private (requires authentication)
	mux.Handle("/create_event", protect(handlers.HandleCreateEvent))
	mux.Handle("/delete_event", protect(handlers.HandleDeleteEvent))
	mux.Handle("/change_availability", protect(handlers.HandleChangeAvailability))
	mux.Handle("/get_all_current_events", protect(handlers.HandleGetAllCurrentEvents))

	log.Println("Server is running at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
