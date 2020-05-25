package main

import (
	"github.com/code-and-chill/iskandar/config"
	"github.com/code-and-chill/iskandar/handler"
	"github.com/code-and-chill/iskandar/infra"
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

	pgConf := config.DBConfig{
		Port:     5432,
		Database: "transport",
		Host:     "localhost",
		Username: "application",
		Password: "application",
	}

	postgresClient := infra.PgConnect(pgConf, log)

	defer func() {
		infra.PgDisconnect(postgresClient)
	}()

	bookingAccessor := postgres.NewBookingSchema(postgresClient)
	paymentAccessor := postgres.NewPaymentSchema(postgresClient)
	ticketAccessor := postgres.NewTicketSchema(postgresClient)
	invoiceAccessor := postgres.NewInvoiceSchema(postgresClient)

	bookingSvc := service.NewBookingService(bookingAccessor, log)
	paymentSvc := service.NewPaymentService(paymentAccessor, log)
	ticketSvc := service.NewTicketService(ticketAccessor, log)
	invoiceSvc := service.NewInvoiceService(invoiceAccessor, log)

	interactor := handler.NewPresenter(server, bookingSvc, paymentSvc, ticketSvc, invoiceSvc)

	interactor.Booking()
	interactor.Payment()
	interactor.Ticket()
	interactor.Invoice()

	err := server.Run(":8080")
	if err != nil {
		panic("Unable to start service")
	}
}
