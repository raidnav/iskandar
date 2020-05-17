package handlers

import (
	"github.com/code-and-chill/iskandar/repositories/models"
	"github.com/code-and-chill/iskandar/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type BookingHandler struct {
	bookingSvc services.BookingService
	log        logrus.FieldLogger
}

type BookingController interface {
	Book() gin.HandlerFunc
	Fetch() gin.HandlerFunc
	Modify() gin.HandlerFunc
	Cancel() gin.HandlerFunc
}

func NewBookingHandler(bookingSvc services.BookingService, log logrus.FieldLogger) BookingController {
	return &BookingHandler{
		bookingSvc: bookingSvc,
		log:        log,
	}
}

func (b BookingHandler) Book() gin.HandlerFunc {
	return func(context *gin.Context) {
		var bookingData models.Booking
		parseError := context.BindJSON(&bookingData)
		if parseError != nil {
			b.log.Warn("Error when parsing json: ", parseError.Error())
		}
		err := b.bookingSvc.Book(bookingData)
		if err != nil {
			context.JSON(500, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "There was an unexpected error during creating a booking",
			})
		} else {
			context.JSON(200, gin.H{
				"code":    http.StatusCreated,
				"message": "Booking Successfully created",
			})
		}
	}
}

func (b BookingHandler) Fetch() gin.HandlerFunc {
	return func(context *gin.Context) {
		bookingId, exist := context.Params.Get("id")
		if !exist {
			context.JSON(400, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Id not provided",
			})
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				b.log.Warn("Unable to convert bookingId: " + bookingId)
			}

			booking, err := b.bookingSvc.Fetch(id)

			if err != nil {
				switch err.Error() {
				case "NOT_FOUND":
					context.JSON(404, gin.H{
						"code": http.NotFound,
						"data": nil,
					})
				default:
					context.JSON(500, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "There was an unexpected error during getting booking data",
					})
				}
			} else {
				context.JSON(200, gin.H{
					"code": http.StatusOK,
					"data": booking,
				})
			}
		}
	}
}

func (b BookingHandler) Modify() gin.HandlerFunc {
	return func(context *gin.Context) {
		bookingId, idExist := context.Params.Get("id")
		status, statusExist := context.Params.Get("status")
		if !(idExist || statusExist) {
			context.JSON(400, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Either booking id or status is not provided",
			})
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				b.log.Warn("Unable to convert bookingId: " + bookingId)
			}

			booking, err := b.bookingSvc.Modify(id, status)

			if err != nil {
				context.JSON(500, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "There was an unexpected error during getting booking data",
				})
			} else {
				context.JSON(200, gin.H{
					"code": http.StatusOK,
					"data": booking,
				})
			}
		}
	}
}

func (b BookingHandler) Cancel() gin.HandlerFunc {
	return func(context *gin.Context) {
		bookingId, exist := context.Params.Get("id")
		reason, reasonExist := context.Params.Get("reason")
		if !(exist || reasonExist) {
			context.JSON(400, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Either booking id or reason is not provided",
			})
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				b.log.Warn("Unable to convert bookingId: " + bookingId)
			}

			booking, err := b.bookingSvc.Cancel(id, reason)

			if err != nil {
				switch err.Error() {
				case "NOT_FOUND":
					context.JSON(404, gin.H{
						"code": http.NotFound,
						"data": nil,
					})
				default:
					context.JSON(500, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "There was an unexpected error during getting booking data",
					})
				}
			} else {
				context.JSON(200, gin.H{
					"code": http.StatusOK,
					"data": booking,
				})
			}
		}
	}
}
