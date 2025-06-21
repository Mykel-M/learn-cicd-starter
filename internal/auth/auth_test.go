package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test 1: Valid API key
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "ApiKey abc123")

	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if apiKey != "abc123" {
		t.Errorf("expected 'abc123', got '%s'", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	// Test 2: Missing Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	_, err := GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
