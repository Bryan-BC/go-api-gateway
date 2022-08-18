package main

import (
	"log"

	"github.com/Bryan-BC/go-api-gateway/pkg/auth"
	"github.com/Bryan-BC/go-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Panicf("Error loading config file, %s", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)

	log.Printf("%+v", authSvc)

	r.Run(c.Port)
}
