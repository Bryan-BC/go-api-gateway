package course

import (
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/Bryan-BC/go-api-gateway/pkg/course/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CourseServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.CourseSvcURL, grpc.WithInsecure())

	if err != nil {
		log.Panicf("Error connecting to course service, %s", err)
	}

	return &ServiceClient{Client: pb.NewCourseServiceClient(cc)}
}
