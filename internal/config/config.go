package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	HTTPAddr    string
}

func MustLoad() Config {
	_ = godotenv.Load()

	dbURL := getenv("DATABASE_URL", "")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return Config{
		DatabaseURL: dbURL,
		HTTPAddr:    getenv("HTTP_ADDR", ":8080"),
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
