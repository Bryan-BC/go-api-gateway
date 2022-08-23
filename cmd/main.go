package main

import (
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/auth"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/Bryan-BC/go-api-gateway/pkg/course"
	"github.com/Bryan-BC/go-api-gateway/pkg/timetable"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Panicf("Error loading config file, %s", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	course.RegisterRoutes(r, &c, &authSvc)
	timetable.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
