package timetable

import (
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/Bryan-BC/go-api-gateway/pkg/timetable/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.TimetableServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.TimetableSvcURL, grpc.WithInsecure())
	if err != nil {
		log.Panicf("Error connecting to timetable service, %s", err)
	}

	return &ServiceClient{Client: pb.NewTimetableServiceClient(cc)}
}
