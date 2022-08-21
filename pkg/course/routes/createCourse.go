package course

import (
	"context"
	"log"
	"net/http"

	"github.com/Bryan-BC/go-api-gateway/pkg/course/pb"
	"github.com/gin-gonic/gin"
)

type CreateCourseRequest struct {
	Course *pb.Course `json:"course"`
}

func CreateCourse(ctx *gin.Context, courseSvc pb.CourseServiceClient) {
	var req CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicf("Error binding json, %s", err)
		return
	}

	resp, err := courseSvc.CreateCourse(context.Background(), &pb.CreateCourseRequest{
		Course: req.Course,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Panicf("Error calling course service, %s", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
