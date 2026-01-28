package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_PropagatesRequestID(t *testing.T) {
	expectedRequestID := "test-request-id-123"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request ID was propagated
		gotRequestID := r.Header.Get(RequestIDKey)
		assert.Equal(t, expectedRequestID, gotRequestID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	// Create context with request ID
	ctx := context.WithValue(context.Background(), RequestIDKey, expectedRequestID)

	resp, err := client.Get(ctx, "/test")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

func TestClient_PostWithBody(t *testing.T) {
	type testBody struct {
		Name string `json:"name"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, ContentTypeJSON, r.Header.Get("Content-Type"))

		var body testBody
		json.NewDecoder(r.Body).Decode(&body)
		assert.Equal(t, "test", body.Name)

		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	resp, err := client.Post(context.Background(), "/test", testBody{Name: "test"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	resp.Body.Close()
}

func TestClient_WithAuth(t *testing.T) {
	expectedToken := "my-jwt-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		assert.Equal(t, "Bearer "+expectedToken, authHeader)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	resp, err := client.Get(context.Background(), "/protected", WithAuth(expectedToken))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

func TestClient_WithTimeout(t *testing.T) {
	client := New(WithTimeout(5 * time.Second))
	assert.Equal(t, 5*time.Second, client.httpClient.Timeout)
}

func TestClient_WithCustomHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "custom-value", r.Header.Get("X-Custom-Header"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	resp, err := client.Get(context.Background(), "/test", WithHeader("X-Custom-Header", "custom-value"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

func TestGetRequestIDFromContext(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		wantID    string
	}{
		{
			name:   "With request ID",
			ctx:    context.WithValue(context.Background(), RequestIDKey, "test-id"),
			wantID: "test-id",
		},
		{
			name:   "Without request ID",
			ctx:    context.Background(),
			wantID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRequestIDFromContext(tt.ctx)
			assert.Equal(t, tt.wantID, got)
		})
	}
}
