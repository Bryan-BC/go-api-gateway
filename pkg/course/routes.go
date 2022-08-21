package course

import (
	"github.com/Bryan-BC/go-api-gateway/pkg/auth"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	routes "github.com/Bryan-BC/go-api-gateway/pkg/course/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ctx *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	svc := NewServiceClient(c)
	a := auth.NewAuthMiddlewareConfig(authSvc)

	routes := ctx.Group("/course")
	routes.Use(a.Authentication)
	routes.POST("/", svc.CreateCourse)
	routes.GET("/:id", svc.GetCourse)
	routes.DELETE("/:id", svc.DeleteCourse)
}

func (svc *ServiceClient) CreateCourse(ctx *gin.Context) {
	routes.CreateCourse(ctx, svc.Client)
}

func (svc *ServiceClient) GetCourse(ctx *gin.Context) {
	routes.GetCourse(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteCourse(ctx *gin.Context) {
	routes.DeleteCourse(ctx, svc.Client)
}
