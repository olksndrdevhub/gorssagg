package auth

import (
	"errors"
	"net/http"
	"strings"
)

// get the API key from the headers of the request
// Example:
// Authorization: ApiKey {API_KEY}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("mallformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("mallformed first part of Authorization header")
	}
	return vals[1], nil
}
