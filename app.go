package main

import (
	"github.com/code-and-chill/iskandar/handler"
	"github.com/code-and-chill/iskandar/infrastructure"
	"github.com/code-and-chill/iskandar/middleware"
	"github.com/code-and-chill/iskandar/repository/postgres"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	server := gin.New()
	server.Use(middleware.Logger(log), gin.Recovery())

	dbConf := infrastructure.DbConfig{
		Port:     5432,
		Database: "transport",
		Host:     "localhost",
		Username: "application",
		Password: "application",
	}

	db := infrastructure.Connect(dbConf, log)
	defer infrastructure.DisConnect(db)

	bookingAccessor := postgres.NewBookingSchema(db)
	paymentAccessor := postgres.NewPaymentSchema(db)

	bookingSvc := service.NewBookingService(bookingAccessor, log)
	paymentSvc := service.NewPaymentService(paymentAccessor, log)

	booking := server.Group("/booking")
	{
		bkHandler := handler.NewBookingHandler(bookingSvc)
		booking.POST("", bkHandler.Book())
		booking.GET("/", bkHandler.Fetch())
		booking.PUT("/", bkHandler.Modify())
		booking.DELETE("/", bkHandler.Cancel())
	}

	payment := server.Group("/payment")
	{
		pgHandler := handler.NewPaymentHandler(paymentSvc)
		payment.GET("/", pgHandler.GenerateRequestSpec())
		payment.POST("", pgHandler.Pay())
		payment.DELETE("/", pgHandler.Cancel())
	}

	err := server.Run(":8080")
	if err != nil {
		panic("Unable to start service")
	}
}
