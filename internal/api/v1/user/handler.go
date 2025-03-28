package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:id", h.Return())
	}
}

func (h *Handler) Return() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id := c.Param("id")
		//user, err := userService.GetUserByID(id)
		//if err != nil {
		//	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		//	return
		//}
		c.JSON(http.StatusOK, "")
	}
}
