package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wow/nigger/internal/dto"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Normal string", "Hello", "olleH"},
		{"Empty string", "", ""},
		{"Single char", "a", "a"},
		{"Palindrome", "racecar", "racecar"},
		{"With spaces", "Hello World", "dlroW olleH"},
		{"Unicode", "你好", "好你"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEcho(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       interface{}
		wantStatus int
		wantCode   string
	}{
		{
			name:       "Valid request",
			body:       dto.EchoRequest{Message: "Hello"},
			wantStatus: http.StatusOK,
			wantCode:   "SUCCESS",
		},
		{
			name:       "Empty message",
			body:       map[string]string{"message": ""},
			wantStatus: http.StatusBadRequest,
			wantCode:   "VALIDATION_ERROR",
		},
		{
			name:       "Missing message field",
			body:       map[string]string{},
			wantStatus: http.StatusBadRequest,
			wantCode:   "VALIDATION_ERROR",
		},
		{
			name:       "Invalid JSON",
			body:       "not json",
			wantStatus: http.StatusBadRequest,
			wantCode:   "VALIDATION_ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.POST("/echo", Echo)

			var bodyBytes []byte
			if str, ok := tt.body.(string); ok {
				bodyBytes = []byte(str)
			} else {
				bodyBytes, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var resp map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Equal(t, tt.wantCode, resp["code"])
		})
	}
}
