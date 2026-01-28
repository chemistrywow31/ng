package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		existingID     string
		expectGenerate bool
	}{
		{
			name:           "Generate new ID when none provided",
			existingID:     "",
			expectGenerate: true,
		},
		{
			name:           "Use existing ID when provided",
			existingID:     "existing-request-id-123",
			expectGenerate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)

			r.Use(RequestID())
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"request_id": GetRequestID(c)})
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.existingID != "" {
				req.Header.Set(RequestIDKey, tt.existingID)
			}
			c.Request = req

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			responseID := w.Header().Get(RequestIDKey)
			assert.NotEmpty(t, responseID)

			if tt.expectGenerate {
				assert.Len(t, responseID, 36) // UUID format
			} else {
				assert.Equal(t, tt.existingID, responseID)
			}
		})
	}
}
