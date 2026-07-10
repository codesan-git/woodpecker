package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var body HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", body.Status)
	}

	if body.Version != "dev" {
		t.Errorf("expected version 'dev', got '%s'", body.Version)
	}
}

func TestHealthHandlerVersionEnv(t *testing.T) {
	t.Setenv("APP_VERSION", "1.0.0")

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	var body HealthResponse
	json.NewDecoder(w.Result().Body).Decode(&body)

	if body.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got '%s'", body.Version)
	}
}

func TestGetVersionDefault(t *testing.T) {
	if v := getVersion(); v != "dev" {
		t.Errorf("expected default version 'dev', got '%s'", v)
	}
}

func TestGetVersionFromEnv(t *testing.T) {
	t.Setenv("APP_VERSION", "2.0.0")
	if v := getVersion(); v != "2.0.0" {
		t.Errorf("expected version '2.0.0', got '%s'", v)
	}
}
