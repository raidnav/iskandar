package main

import (
	"github.com/code-and-chill/iskandar/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	server := gin.New()
	server.Use(middlewares.Logger(log), gin.Recovery())

	booking := server.Group("/booking")
	{
		booking.POST("", func(context *gin.Context) {})
		booking.GET("/", func(context *gin.Context) {})
	}

	err := server.Run(":8080")
	if err != nil {
		panic("Unable to start service")
	}
}
