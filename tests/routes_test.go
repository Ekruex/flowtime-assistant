package tests

import (
	"flowtime-backend/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Helper function to create a test router
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", routes.HealthCheck)
	r.POST("/session/start", routes.StartSession)
	r.POST("/session/end", routes.EndSession)
	return r
}

// Test the /health route
func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedBody := `{"status":"OK"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, w.Body.String())
	}
}

// Test the /session/start route
func TestStartSession(t *testing.T) {
	router := setupRouter()

	// Create a test HTTP request
	req, _ := http.NewRequest("POST", "/session/start", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedBody := `{"message":"Session started successfully!"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, w.Body.String())
	}
}

// Test the /session/end route
func TestEndSession(t *testing.T) {
	router := setupRouter()

	// Create a test HTTP request
	req, _ := http.NewRequest("POST", "/session/end", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedBody := `{"message":"Session ended successfully!"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, w.Body.String())
	}
}
