package main

import (
	"log"
	env "spotify-api/src/config"
	album "spotify-api/src/service/album"
)

func main() {
	envConfig := env.LoadEnvConfig()

	log.Printf("Spotify Client ID: %s", envConfig.SpotifyClientID)

	response, err := album.GetData("2cWBwpqMsDJC1ZUwz813lo")
	if err != nil {
		log.Fatalf("failed to get album data: %v", err)
	}

	log.Fatal("Album: ", response)
}
