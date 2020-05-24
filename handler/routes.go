package handler

import (
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
)

type Routes interface {
	Booking(server *gin.Engine) *gin.RouterGroup
	Payment(server *gin.Engine) *gin.RouterGroup
}

type Service struct {
	book service.BookingService
	pay  service.PaymentService
}

func (s Service) Booking(server *gin.Engine) *gin.RouterGroup {
	booking := server.Group("/booking")
	{
		resolver := NewBookingHandler(s.book)
		booking.POST("", func(context *gin.Context) { resolver.Book(context) })
		booking.GET("/", func(context *gin.Context) { resolver.Fetch(context) })
		booking.PUT("/", func(context *gin.Context) { resolver.Modify(context) })
		booking.DELETE("/", func(context *gin.Context) { resolver.Cancel(context) })
	}
	return booking
}

func (s Service) Payment(server *gin.Engine) *gin.RouterGroup {
	payment := server.Group("/payment")
	{
		resolver := NewPaymentHandler(s.pay)
		payment.GET("/", func(context *gin.Context) { resolver.GenerateRequestSpec(context) })
		payment.POST("", func(context *gin.Context) { resolver.Pay(context) })
		payment.DELETE("/", func(context *gin.Context) { resolver.Cancel(context) })
	}
	return payment
}

func NewService(booking service.BookingService, payment service.PaymentService) Routes {
	return &Service{
		book: booking,
		pay:  payment,
	}
}
