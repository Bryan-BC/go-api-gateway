package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Bryan-BC/go-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func NewAuthMiddlewareConfig(svc *ServiceClient) *AuthMiddlewareConfig {
	return &AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) Authentication(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		log.Panicln("Authorization header is empty. Status: 401")
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) != 2 {
		log.Panicln("Authorization header is invalid. Status: 401")
		return
	}

	resp, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || resp.Status != http.StatusOK {
		log.Panicln("Authorization header is invalid. Status: 401")
	}

	ctx.Next()
}
