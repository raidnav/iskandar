package main

import (
	"github.com/code-and-chill/iskandar/handlers"
	"github.com/code-and-chill/iskandar/infrastructures"
	"github.com/code-and-chill/iskandar/middlewares"
	"github.com/code-and-chill/iskandar/repositories/postgres"
	"github.com/code-and-chill/iskandar/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	server := gin.New()
	server.Use(middlewares.Logger(log), gin.Recovery())

	dbConf := infrastructures.DbConfig{
		Port:     5432,
		Database: "transport",
		Host:     "localhost",
		Username: "application",
		Password: "application",
	}

	db := infrastructures.Connect(dbConf, log)
	defer infrastructures.DisConnect(db)

	bookingAccessor := postgres.NewPostgresBookingSchema(db, log)

	bookingSvc := services.NewBookingService(bookingAccessor, log)

	booking := server.Group("/booking")
	{
		handler := handlers.NewBookingHandler(bookingSvc, log)
		booking.POST("", handler.Book())
		booking.GET("/", handler.Fetch())
		booking.PUT("/", handler.Modify())
		booking.DELETE("/", handler.Cancel())
	}

	err := server.Run(":8080")
	if err != nil {
		panic("Unable to start service")
	}
}
