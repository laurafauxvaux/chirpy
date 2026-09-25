package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	keyData := headers.Get("Authorization")
	if keyData == "" {
		return "", fmt.Errorf("authorization header doesn't exist")
	}

	if !strings.HasPrefix(keyData, "ApiKey ") {
		return "", fmt.Errorf("no API key in authorization header")
	}

	apiKey := strings.TrimSpace(strings.TrimPrefix(keyData, "ApiKey "))

	return apiKey, nil
}
