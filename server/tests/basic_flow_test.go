package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_server/internal/config"
	"go_server/internal/repository"
	"go_server/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func createAuthenticatedUser(t *testing.T, db *sql.DB) (string, string, string) {
	email := fmt.Sprintf("user.%d@iaeste.pl", time.Now().UnixNano())
	password := "testpassword"
	firstName, lastName, err := repository.ExtractNameFromEmail(email)
	if err != nil {
		t.Fatalf("Invalid email format: %s", err)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err = db.Exec(`INSERT INTO accounts (email, password_hash, verification_token, is_verified)
		VALUES ($1, $2, '', TRUE)`, email, string(hashed))
	if err != nil {
		t.Fatalf("Helper: Błąd insert accounts: %v", err)
	}

	var userID int
	err = db.QueryRow(`INSERT INTO users (first_name, last_name, email, "group")
		VALUES ($2, $3, $1, 'IT') RETURNING id`, email, firstName, lastName).Scan(&userID)
	if err != nil {
		t.Fatalf("Helper: Błąd insert users: %v", err)
	}

	claims := &middleware.Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenObj.SignedString([]byte(config.JWTSecret))
	if err != nil {
		t.Fatalf("Helper: Błąd podpisywania tokena: %v", err)
	}

	return tokenString, email, password
}

func TestLoginFlow(t *testing.T) {
	testDB := setupTestDB(t)
	router := setupRouter()

	token, email, password := createAuthenticatedUser(t, testDB)

	loginPayload, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	reqLogin := httptest.NewRequest("POST", "/login", bytes.NewBuffer(loginPayload))
	reqLogin.Header.Set("Content-Type", "application/json")
	wLogin := httptest.NewRecorder()

	router.ServeHTTP(wLogin, reqLogin)

	if wLogin.Code != http.StatusAccepted {
		t.Fatalf("Login failed. Code: %d, Body: %s", wLogin.Code, wLogin.Body.String())
	}

	var resp map[string]string
	if err := json.NewDecoder(wLogin.Body).Decode(&resp); err != nil {
		t.Fatalf("Response decoding error: %v", err)
	}

	token, ok := resp["token"]
	if !ok || token == "" {
		t.Fatal("No token in login response")
	}

	t.Logf("Success: logged in and received a token: %s...", token[:10])
}
