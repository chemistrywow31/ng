package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wow/nigger/internal/dto"
	"github.com/wow/nigger/internal/middleware"
)

// Login godoc
// @Summary      Login
// @Description  Authenticate and get JWT token (test endpoint: use admin/admin123)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "Login credentials"
// @Success      200      {object}  dto.Response{data=dto.LoginResponse}
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      401      {object}  dto.ErrorResponse
// @Router       /api/v1/auth/login [post]
func Login(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		// Test credentials: admin/admin123
		if req.Username != "admin" || req.Password != "admin123" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    "UNAUTHORIZED",
				Message: "Invalid credentials",
				Data:    nil,
			})
			return
		}

		expiresAt := time.Now().Add(72 * time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "user-123",
			"exp": expiresAt.Unix(),
			"iat": time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to generate token",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, dto.Response{
			Code:    "SUCCESS",
			Message: "Login successful",
			Data: dto.LoginResponse{
				Token:     tokenString,
				ExpiresAt: expiresAt.UTC().Format(time.RFC3339),
			},
		})
	}
}

// GetMe godoc
// @Summary      Get current user
// @Description  Get current authenticated user info
// @Tags         auth
// @Produce      json
// @Success      200  {object}  dto.Response{data=dto.UserInfoResponse}
// @Failure      401  {object}  dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/auth/me [get]
func GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, dto.Response{
		Code:    "SUCCESS",
		Message: "User info retrieved",
		Data: dto.UserInfoResponse{
			UserID:    userID.(string),
			RequestID: middleware.GetRequestID(c),
		},
	})
}
