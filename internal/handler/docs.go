package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wow/nigger/internal/dto"
)

// GetDocs godoc
// @Summary      Get user manual
// @Description  Returns the user manual markdown content
// @Tags         system
// @Produce      json
// @Success      200  {object}  dto.DocsResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Router       /api/docs [get]
func GetDocs(c *gin.Context) {
	content, err := os.ReadFile("docs/public/user-manual.md")
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "NOT_FOUND",
			Message: "Documentation not found",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, dto.DocsResponse{
		Content: string(content),
	})
}
