package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wow/nigger/internal/dto"
	"github.com/wow/nigger/internal/middleware"
)

// Ping godoc
// @Summary      Ping test
// @Description  Simple ping endpoint to test connectivity
// @Tags         test
// @Produce      json
// @Success      200  {object}  dto.Response{data=dto.PingResponse}
// @Security     BearerAuth
// @Router       /api/v1/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{
		Code:    "SUCCESS",
		Message: "pong",
		Data: dto.PingResponse{
			Message:   "pong",
			RequestID: middleware.GetRequestID(c),
		},
	})
}
