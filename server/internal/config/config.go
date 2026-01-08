package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseURL   string
	EmailUser     string
	EmailPassword string
	JWTSecret     string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	EmailUser = os.Getenv("EMAIL_USER")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	JWTSecret = os.Getenv("JWT_SECRET")

	log.Println("Configuration loaded")
}
