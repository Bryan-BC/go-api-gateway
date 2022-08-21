package course

import (
	"context"
	"log"
	"net/http"

	"github.com/Bryan-BC/go-api-gateway/pkg/course/pb"
	"github.com/gin-gonic/gin"
)

func DeleteCourse(ctx *gin.Context, courseSvc pb.CourseServiceClient) {
	id := ctx.Param("id")

	resp, err := courseSvc.DeleteCourse(context.Background(), &pb.DeleteCourseRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Panicf("Error calling course service, %s", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
