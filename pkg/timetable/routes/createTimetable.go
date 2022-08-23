package timetable

import (
	"context"
	"log"
	"net/http"

	"github.com/Bryan-BC/go-api-gateway/pkg/timetable/pb"
	"github.com/gin-gonic/gin"
)

type CreateTimetableRequest struct {
	Courses []string `json:"courses"`
}

func CreateTimetable(ctx *gin.Context, timetableSvc pb.TimetableServiceClient) {
	var req CreateTimetableRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicf("Error binding json, %s \n", err)
		return
	}

	resp, err := timetableSvc.CreateTimetable(context.Background(), &pb.CreateTimetableRequest{
		Courses: req.Courses,
		UserId:  ctx.GetInt64("userId"),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Panicf("Error calling timetable service, %s \n", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
