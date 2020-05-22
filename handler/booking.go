package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookingHandler struct {
	bookingSvc service.BookingService
}

type BookingController interface {
	Book() gin.HandlerFunc
	Fetch() gin.HandlerFunc
	Modify() gin.HandlerFunc
	Cancel() gin.HandlerFunc
}

func NewBookingHandler(bookingSvc service.BookingService) BookingController {
	return &BookingHandler{
		bookingSvc: bookingSvc,
	}
}

func (b BookingHandler) Book() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		var bookingData models.Booking
		parseError := context.BindJSON(&bookingData)
		if parseError != nil {
			w.Message(http.StatusBadRequest, "Invalid booking spec")
		}
		err := b.bookingSvc.Book(bookingData)
		if err != nil {
			w.Message(http.StatusInternalServerError, "There was an unexpected error during creating a booking")
		} else {
			w.Data(http.StatusOK, nil)
		}
	}
}

func (b BookingHandler) Fetch() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		bookingId, exist := context.Params.Get("id")
		if !exist {
			w.Message(http.StatusBadRequest, "Id not provided")
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
			}

			booking, err := b.bookingSvc.Fetch(id)

			if err != nil {
				switch err.Error() {
				case constant.DbNotFound:
					w.Message(http.StatusNotFound, "Booking Id was not found")
				default:
					w.Message(http.StatusInternalServerError, "There was an unexpected error during getting booking data")
				}
			} else {
				w.Data(http.StatusOK, booking)
			}
		}
	}
}

func (b BookingHandler) Modify() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		bookingId, idExist := context.Params.Get("id")
		status, statusExist := context.Params.Get("status")
		if !(idExist || statusExist) {
			w.Message(http.StatusBadRequest, "Either booking id or status is not provided")
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
			}

			err := b.bookingSvc.Modify(id, status)
			if err != nil {
				w.Message(http.StatusInternalServerError, "There was an unexpected error during getting booking data")
			} else {
				w.Data(http.StatusOK, nil)
			}
		}
	}
}

func (b BookingHandler) Cancel() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		bookingId, exist := context.Params.Get("id")
		reason, reasonExist := context.Params.Get("reason")
		if !(exist || reasonExist) {
			w.Message(http.StatusBadRequest, "Either booking id or reason is not provided")
		} else {
			id, numError := strconv.Atoi(bookingId)
			if numError != nil {
				w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
			}

			err := b.bookingSvc.Cancel(id, reason)

			if err != nil {
				switch err.Error() {
				case constant.DbNotFound:
					w.Message(http.StatusNotFound, "Requested booking id was not found")
				default:
					w.Message(http.StatusInternalServerError, "There was an unexpected error during cancelling booking")
				}
			} else {
				w.Data(http.StatusNoContent, nil)
			}
		}
	}
}
