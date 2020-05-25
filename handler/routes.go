package handler

import (
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
)

type Routes interface {
	Booking() *gin.RouterGroup
	Payment() *gin.RouterGroup
	Ticket() *gin.RouterGroup
	Invoice() *gin.RouterGroup
}

type Service struct {
	presenter  *gin.Engine
	bookingSvc service.BookingService
	paymentSvc service.PaymentService
	ticketSvc  service.TicketService
	invoiceSvc service.InvoiceService
}

func (s Service) Ticket() *gin.RouterGroup {
	ticket := s.presenter.Group("/ticket")
	{
		resolver := NewTicketHandler(s.ticketSvc, s.bookingSvc)
		ticket.GET("/", func(context *gin.Context) { resolver.Fetch(context) })
		ticket.POST("/", func(context *gin.Context) { resolver.Save(context) })
	}
	return ticket
}

func (s Service) Invoice() *gin.RouterGroup {
	invoice := s.presenter.Group("/invoice")
	{
		resolver := NewInvoiceHandler(s.invoiceSvc)
		invoice.GET("/", func(context *gin.Context) { resolver.Fetch(context) })
		invoice.POST("/", func(context *gin.Context) { resolver.Save(context) })
	}
	return invoice
}

func (s Service) Booking() *gin.RouterGroup {
	booking := s.presenter.Group("/booking")
	{
		resolver := NewBookingHandler(s.bookingSvc)
		booking.POST("", func(context *gin.Context) { resolver.Book(context) })
		booking.GET("/:id", func(context *gin.Context) { resolver.Fetch(context) })
		booking.PUT("/:id", func(context *gin.Context) { resolver.Modify(context) })
		booking.DELETE("/:id", func(context *gin.Context) { resolver.Cancel(context) })
	}
	return booking
}

func (s Service) Payment() *gin.RouterGroup {
	payment := s.presenter.Group("/payment")
	{
		resolver := NewPaymentHandler(s.paymentSvc)
		payment.GET("/", func(context *gin.Context) { resolver.GenerateRequestSpec(context) })
		payment.POST("", func(context *gin.Context) { resolver.Pay(context) })
		payment.DELETE("/", func(context *gin.Context) { resolver.Cancel(context) })
	}
	return payment
}

func NewPresenter(server *gin.Engine,
	booking service.BookingService,
	payment service.PaymentService,
	ticket service.TicketService,
	invoice service.InvoiceService) Routes {
	return &Service{
		presenter:  server,
		bookingSvc: booking,
		paymentSvc: payment,
		ticketSvc:  ticket,
		invoiceSvc: invoice,
	}
}
