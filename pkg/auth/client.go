package auth

import (
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/auth/pb"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.AuthSvcURL, grpc.WithInsecure())
	if err != nil {
		log.Panicf("Error connecting to auth service, %s", err)
	}

	return &ServiceClient{Client: pb.NewAuthServiceClient(cc)}
}
