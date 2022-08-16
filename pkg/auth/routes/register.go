package auth

import (
	"context"
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, authSvc pb.AuthServiceClient) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Panicf("Error binding json, %s", err)
		return
	}

	resp, err := authSvc.Register(context.Background(), &pb.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		log.Panicf("Error calling auth service, %s", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
