package auth

import (
	routes "github.com/Bryan-BC/go-api-gateway/pkg/auth/routes"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := NewServiceClient(c)

	routes := r.Group("/auth")
	routes.POST("/login", svc.Login)
	routes.POST("/register", svc.Register)

	return svc
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
