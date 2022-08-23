package timetable

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Bryan-BC/go-api-gateway/pkg/timetable/pb"
	"github.com/gin-gonic/gin"
)

func GetTimetable(ctx *gin.Context, timetableSvc pb.TimetableServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panicf("Error parsing id, %s \n", err)
	}

	resp, err := timetableSvc.GetTimetable(context.Background(), &pb.GetTimetableRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Panicf("Error calling timetable service, %s", err)
		return
	}

	ctx.JSON(int(resp.Status), resp)
}
