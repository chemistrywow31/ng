package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wow/nigger/internal/dto"
	"github.com/wow/nigger/internal/middleware"
)

// Echo godoc
// @Summary      Echo message
// @Description  Echoes back the message with additional info
// @Tags         test
// @Accept       json
// @Produce      json
// @Param        request  body      dto.EchoRequest  true  "Echo request"
// @Success      200      {object}  dto.Response{data=dto.EchoResponse}
// @Failure      400      {object}  dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/echo [post]
func Echo(c *gin.Context) {
	var req dto.EchoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    "SUCCESS",
		Message: "Echo successful",
		Data: dto.EchoResponse{
			Original:  req.Message,
			Reversed:  reverseString(req.Message),
			Length:    len(req.Message),
			RequestID: middleware.GetRequestID(c),
		},
	})
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
