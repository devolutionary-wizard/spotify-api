package apirequest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	env "spotify-api/src/config"
	"strings"
)

func MakeRequest(method, url, body string) {
	envConfig := env.LoadEnvConfig()
	req, err := http.NewRequest(method, url, strings.NewReader(body))

	retryCount := 0
	const retryLimit = 1

	for retryCount < retryLimit {
		if err != nil {
			log.Fatalf("Failed to create request: %v", err)
		}

		// Setting the Authorization header with the Bearer token
		req.Header.Set("Authorization", "Bearer "+envConfig.SpotifyACcessToken)

		resq, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer resq.Body.Close()

		if resq.StatusCode != http.StatusOK {
			GetToken()
			log.Fatalf("Error response: %s", resq.Status)
		}

		responseBody, err := io.ReadAll(resq.Body)

		if err != nil {
			log.Fatalf("Error reading response: %v", err)
		}

		fmt.Println(string(responseBody))

		retryCount++
	}
}
