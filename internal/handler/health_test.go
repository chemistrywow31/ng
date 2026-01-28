package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wow/nigger/internal/dto"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name        string
		version     string
		wantStatus  int
		wantVersion string
	}{
		{
			name:        "Returns health status with version",
			version:     "1.0.0",
			wantStatus:  http.StatusOK,
			wantVersion: "1.0.0",
		},
		{
			name:        "Returns health status with different version",
			version:     "2.0.0-beta",
			wantStatus:  http.StatusOK,
			wantVersion: "2.0.0-beta",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.GET("/health", Health(tt.version))

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var resp dto.HealthResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, "ok", resp.Status)
			assert.Equal(t, tt.wantVersion, resp.Version)
			assert.NotEmpty(t, resp.Timestamp)
		})
	}
}
