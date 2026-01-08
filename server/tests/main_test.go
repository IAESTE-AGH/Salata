package tests

import (
	"context"
	"database/sql"
	"go_server/internal/config"
	"go_server/internal/database"
	"go_server/internal/handlers"
	"go_server/middleware"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
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

func setupTestDB(t *testing.T) *sql.DB {
	ctx := context.Background()

	// Create postgres container for testing
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:17"),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %s", err)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	database.DB = db

	t.Cleanup(func() {
		db.Close()
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	schemaPath := filepath.Join("..", "schema.sql")

	content, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("failed to load schema.sql file: %s", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		t.Fatalf("schema.sql execution error: %s", err)
	}

	return db
}

func TestMain(m *testing.M) {
	config.JWTSecret = "test-secret-key"

	code := m.Run()

	os.Exit(code)
}
