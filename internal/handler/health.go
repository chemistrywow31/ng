package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wow/nigger/internal/dto"
)

// Health godoc
// @Summary      Health check
// @Description  Check if the service is running
// @Tags         system
// @Produce      json
// @Success      200  {object}  dto.HealthResponse
// @Router       /health [get]
func Health(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.HealthResponse{
			Status:    "ok",
			Version:   version,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
	}
}
