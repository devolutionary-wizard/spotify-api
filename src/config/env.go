package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var envPath string

func Init() {
	envPath, _ = os.Getwd()
	err := godotenv.Load(filepath.Join(envPath, ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

type Env struct {
	SpotifyAccountBaseURL string
	SpotifyBaseURL        string
	SpotifyClientID       string
	SpotifyClientSecret   string
	SpotifyACcessToken    string
}

func LoadEnvConfig() *Env {
	Init()

	return &Env{
		SpotifyAccountBaseURL: GetEnv("SPOTIFY_ACCOUNTS_BASE_URL"),
		SpotifyBaseURL:        GetEnv("SPOTIFY_BASE_URL"),
		SpotifyClientID:       GetEnv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret:   GetEnv("SPOTIFY_CLIENT_SECRET"),
		SpotifyACcessToken:    GetEnv("SPOTIFY_ACCESS_TOKEN"),
	}
}
