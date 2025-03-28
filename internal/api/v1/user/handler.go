package user

import (
	"github.com/gin-gonic/gin"
	"github.com/majidimanzade/base-go/internal/services/generator"
	"net/http"
)

type Handler struct {
	generator generator.GeneratorService
}

func New(g generator.GeneratorService) *Handler {
	return &Handler{
		generator: g,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/generate", h.Generate())
	}
}

func (h *Handler) Generate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h.generator.Run(); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, "")
	}
}
