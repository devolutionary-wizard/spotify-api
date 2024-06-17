package main

import (
	"log"
	apirequest "spotify-api/src/api-request"
	env "spotify-api/src/config"
)

func main() {
	envConfig := env.LoadEnvConfig()

	log.Printf("Spotify Client ID: %s", envConfig.SpotifyClientID)

	apirequest.MakeRequest("GET", "https://api.spotify.com/v1/albums/4aawyAB9vmqN3uQ7FjRGTy", "")
}
