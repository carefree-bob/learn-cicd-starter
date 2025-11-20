package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey_NoHeader(t *testing.T) {
	h := http.Header{}
	_, err := auth.GetAPIKey(h)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err != auth.ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_ValidHeader(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey ABC123")

	key, err := auth.GetAPIKey(h)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key != "ABC123" {
		t.Fatalf("expected API key %q, got %q", "ABC123", key)
	}
}
