package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// SplitAuthHeader - Splits auth header into slice of strings
func SplitAuthHeader(authHeader string) ([]string, error) {
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) < 2 || splitAuthHeader[0] != "ApiKey" {
		return nil, errors.New("malformed authorizationHeader")
	}

	return splitAuthHeader, nil
}

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth, err := SplitAuthHeader(authHeader)
	if err != nil {
		return "", err
	}

	return splitAuth[1], nil
}
