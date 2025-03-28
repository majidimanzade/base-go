package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/majidimanzade/base-go/internal/api/v1"
	"github.com/majidimanzade/base-go/internal/services"
)

func InitializeHandlers(r *gin.Engine, di *services.Services) {
	v1.NewV1(di).RegisterRoutes(r)
}
