package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/majidimanzade/base-go/internal/api/v1"
)

func InitializeHandlers(r *gin.Engine) {
	v1.NewV1().RegisterRoutes(r)
}
