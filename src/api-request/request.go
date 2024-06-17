package apirequest

import (
	"fmt"
	"io"
	"net/http"
	env "spotify-api/src/config"
	"strings"
)

func MakeRequest(method, url, body string) (string, error) {
	envConfig := env.LoadEnvConfig()
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	retryCount := 0
	const retryLimit = 1

	for retryCount < retryLimit {
		// Setting the Authorization header with the Bearer token
		req.Header.Set("Authorization", "Bearer "+envConfig.SpotifyACcessToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			GetToken()
			retryCount++
			continue
		}

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response: %v", err)
		}

		return string(responseBody), nil
	}

	return "", fmt.Errorf("request failed after %d retries", retryLimit)
}
