package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetEnvFallback(t *testing.T) {
	key := "TASKFLOW_TEST_UNSET_VAR"
	t.Setenv(key, "")
	if got := getEnv(key, "fallback"); got != "fallback" {
		t.Fatalf("expected fallback value, got %q", got)
	}
}

func TestGetEnvOverride(t *testing.T) {
	key := "TASKFLOW_TEST_SET_VAR"
	t.Setenv(key, "custom")
	if got := getEnv(key, "fallback"); got != "custom" {
		t.Fatalf("expected custom value, got %q", got)
	}
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	setupRoutes(router)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("health response is not valid JSON: %v", err)
	}
	if body["status"] != "healthy" {
		t.Fatalf("expected status 'healthy', got %v", body["status"])
	}
	if body["service"] != "taskflow-backend" {
		t.Fatalf("expected service 'taskflow-backend', got %v", body["service"])
	}
}

func TestBoardsListEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	setupRoutes(router)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/boards", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
