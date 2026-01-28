package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const testSecret = "test-secret"

func generateTestToken(secret string, exp time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user-123",
		"exp": exp.Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		path       string
		authHeader string
		wantStatus int
		wantUserID bool
	}{
		{
			name:       "Public path /health - no auth required",
			path:       "/health",
			authHeader: "",
			wantStatus: http.StatusOK,
			wantUserID: false,
		},
		{
			name:       "Public path /swagger - no auth required",
			path:       "/swagger/index.html",
			authHeader: "",
			wantStatus: http.StatusOK,
			wantUserID: false,
		},
		{
			name:       "Protected path - missing auth header",
			path:       "/api/v1/protected",
			authHeader: "",
			wantStatus: http.StatusUnauthorized,
			wantUserID: false,
		},
		{
			name:       "Protected path - invalid format (no Bearer)",
			path:       "/api/v1/protected",
			authHeader: "invalid-token",
			wantStatus: http.StatusUnauthorized,
			wantUserID: false,
		},
		{
			name:       "Protected path - invalid token",
			path:       "/api/v1/protected",
			authHeader: "Bearer invalid-token",
			wantStatus: http.StatusUnauthorized,
			wantUserID: false,
		},
		{
			name:       "Protected path - expired token",
			path:       "/api/v1/protected",
			authHeader: "Bearer " + generateTestToken(testSecret, time.Now().Add(-1*time.Hour)),
			wantStatus: http.StatusUnauthorized,
			wantUserID: false,
		},
		{
			name:       "Protected path - valid token",
			path:       "/api/v1/protected",
			authHeader: "Bearer " + generateTestToken(testSecret, time.Now().Add(1*time.Hour)),
			wantStatus: http.StatusOK,
			wantUserID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.Use(Auth(testSecret))
			r.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})
			r.GET("/swagger/*any", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})
			r.GET("/api/v1/protected", func(c *gin.Context) {
				userID, exists := c.Get("user_id")
				c.JSON(http.StatusOK, gin.H{"user_id": userID, "exists": exists})
			})

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestIsPublicPath(t *testing.T) {
	tests := []struct {
		path     string
		isPublic bool
	}{
		{"/health", true},
		{"/api/v1/auth/login", true},
		{"/api/docs", true},
		{"/swagger/index.html", true},
		{"/swagger/doc.json", true},
		{"/api/v1/users", false},
		{"/api/v1/orders", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equal(t, tt.isPublic, isPublicPath(tt.path))
		})
	}
}
