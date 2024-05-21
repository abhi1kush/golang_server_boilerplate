package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestAuthenticationMiddleware(t *testing.T) {
// 	// Create a new request with a dummy JWT token in the Authorization header
// 	req := httptest.NewRequest("GET", "/api/users", nil)
// 	req.Header.Set("Authorization", "Bearer valid-token")

// 	// Create a response recorder to record the response
// 	rr := httptest.NewRecorder()

// 	// Define a handler function that just writes a success message
// 	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Authenticated"))
// 	})

// 	// Create a new instance of the authentication middleware
// 	authMiddleware := AuthenticationMiddleware(handler)

// 	// Call the authentication middleware with the request and response recorder
// 	authMiddleware.ServeHTTP(rr, req)

// 	// Check if the status code is 200 OK
// 	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")

// 	// Check if the response body contains the expected message
// 	expected := "Authenticated"
// 	assert.Equal(t, expected, rr.Body.String(), "Response body should contain 'Authenticated'")
// }

func TestAuthenticationMiddleware2(t *testing.T) {
	// Generate a valid JWT token for testing
	validToken, err := GenerateToken("testuser")
	assert.NoError(t, err, "Error generating token")

	// Create a new request with the valid JWT token in the Authorization header
	req := httptest.NewRequest("GET", "/api/example", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Define a handler function that just writes a success message
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	})

	// Create a new instance of the authentication middleware
	authMiddleware := AuthenticationMiddleware(handler)

	// Call the authentication middleware with the request and response recorder
	authMiddleware.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")

	// Check if the response body contains the expected message
	expected := "Authenticated"
	assert.Equal(t, expected, rr.Body.String(), "Response body should contain 'Authenticated'")
}
