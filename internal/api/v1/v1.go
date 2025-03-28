package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/majidimanzade/base-go/internal/api/v1/user"
	"github.com/majidimanzade/base-go/internal/services"
)

type V1 struct {
	UserHandler *user.Handler
}

func NewV1(di *services.Services) *V1 {
	return &V1{
		UserHandler: user.New(di.Generator),
	}
}

func (v1 *V1) RegisterRoutes(e *gin.Engine) {
	router := e.Group("/v1")
	v1.UserHandler.RegisterRoutes(router)
}
