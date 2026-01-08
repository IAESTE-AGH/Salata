package tests

import (
	"go_server/internal/config"
	"go_server/internal/database"
	"go_server/internal/handlers"
	"go_server/middleware"
	"net/http"
	"os"
	"testing"
)

func setupRouter() *http.ServeMux {
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

	return mux
}

func TestMain(m *testing.M) {
	database.Init(config.DatabaseURL)

	code := m.Run()

	os.Exit(code)
}
