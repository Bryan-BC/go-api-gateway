package course

import (
	"context"
	"log"
	"net/http"

	"github.com/Bryan-BC/go-api-gateway/pkg/course/pb"
	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context, courseSvc pb.CourseServiceClient) {
	id := ctx.Param("id")

	resp, err := courseSvc.GetCourse(context.Background(), &pb.GetCourseRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Panicf("Error calling course service, %s", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
