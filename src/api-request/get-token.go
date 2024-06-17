package apirequest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	env "spotify-api/src/config"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken() *TokenResponse {
	log.Println("Getting token")
	envConfig := env.LoadEnvConfig()
	urlStr := envConfig.SpotifyAccountBaseURL + "/api/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", envConfig.SpotifyClientID)
	data.Set("client_secret", envConfig.SpotifyClientSecret)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response: %s", body)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	StoreTokenInEnv(&tokenResponse)

	return &tokenResponse
}

func StoreTokenInEnv(tokenResponse *TokenResponse) {
	content, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("Failed to read .env file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	var updatedLines []string
	for _, line := range lines {
		if !strings.HasPrefix(line, "SPOTIFY_ACCESS_TOKEN=") &&
			!strings.HasPrefix(line, "SPOTIFY_TOKEN_TYPE=") &&
			!strings.HasPrefix(line, "SPOTIFY_EXPIRES_IN=") {
			updatedLines = append(updatedLines, line)
		}
	}

	updatedLines = append(updatedLines, fmt.Sprintf("SPOTIFY_ACCESS_TOKEN=%s", tokenResponse.AccessToken))
	updatedLines = append(updatedLines, fmt.Sprintf("SPOTIFY_TOKEN_TYPE=%s", tokenResponse.TokenType))
	updatedLines = append(updatedLines, fmt.Sprintf("SPOTIFY_EXPIRES_IN=%d", tokenResponse.ExpiresIn))

	updatedContent := strings.Join(updatedLines, "\n")

	err = os.WriteFile(".env", []byte(updatedContent), 0644) // Changed from ioutil.WriteFile to os.WriteFile
	if err != nil {
		log.Fatalf("Failed to write to .env file: %v", err)
	}
}
