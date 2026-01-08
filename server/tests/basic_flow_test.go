package tests

import (
	"bytes"
	"encoding/json"
	"go_server/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestLoginFlow(t *testing.T) {
	testDB := setupTestDB(t)
	router := setupRouter()

	email := "test.email@iaeste.pl"
	password := "testpassword"
	firstName, lastName, err := repository.ExtractNameFromEmail(email)
	if err != nil {
		t.Fatalf("Invalid email format: %s", err)
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Hash error: %v", err)
	}
	passwordHash := string(hashedBytes)

	_, err = testDB.Exec(`
		INSERT INTO accounts (email, password_hash, verification_token, is_verified)
		VALUES ($1, $2, '', TRUE)`,
		email, passwordHash,
	)
	if err != nil {
		t.Fatalf("Error inserting into accounts: %v", err)
	}

	_, err = testDB.Exec(`
		INSERT INTO users (first_name, last_name, email, "group")
		VALUES ($1, $2, $3, 'IT')`,
		firstName, lastName, email,
	)
	if err != nil {
		t.Fatalf("Error inserting into users: %v", err)
	}

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
