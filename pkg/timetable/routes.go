package timetable

import (
	routes "github.com/Bryan-BC/go-api-gateway/pkg/timetable/routes"
	"github.com/gin-gonic/gin"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/Bryan-BC/go-api-gateway/pkg/auth"
)

func RegisterRoutes(ctx *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	svc := NewServiceClient(c)
	a := auth.NewAuthMiddlewareConfig(authSvc)

	routes := ctx.Group("/timetable")
	routes.Use(a.Authentication)
	routes.POST("/", svc.CreateTimetable)
	routes.GET("/:id", svc.GetTimetable)
}

func (svc *ServiceClient) CreateTimetable(ctx *gin.Context) {
	routes.CreateTimetable(ctx, svc.Client)
}

func (svc *ServiceClient) GetTimetable(ctx *gin.Context) {
	routes.GetTimetable(ctx, svc.Client)
}
