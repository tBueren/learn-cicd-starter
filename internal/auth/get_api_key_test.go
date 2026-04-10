package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func unused() {
	// this funciton is only here to piss the static check off
}

func TestNoAuthHeaderGiven(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestBadAuthHeaderGiven_NoIdentifier(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "random-api-key")

	res, err := GetAPIKey(headers)

	if res != "" || err == nil {
		t.Fatalf("expected error to be thrown. Got res %v, and err val %v", res, err)
	}
}

func TestBadAuthHeaderGiven_WrongIdentifier(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer random-api-key")

	res, err := GetAPIKey(headers)

	if res != "" || err == nil {
		t.Fatalf("expected error to be thrown. Got res %v, and err val %v", res, err)
	}
}

func TestApiKeyReturnedCorrectly(t *testing.T) {
	apiKey := "random-api-key"
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("ApiKey %s", apiKey))

	res, err := GetAPIKey(headers)

	if res != apiKey || err != nil {
		t.Fatalf("Unexpexted error to be thrown. Got res %v, and err val %v, expected val %v", res, err, apiKey)
	}
}
