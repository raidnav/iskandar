package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/handler/helper"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Booking struct {
	bookingSvc service.BookingService
}

type BookingHandler interface {
	Book(context *gin.Context) helper.ResponseWriter
	Fetch(context *gin.Context) helper.ResponseWriter
	Modify(context *gin.Context) helper.ResponseWriter
	Cancel(context *gin.Context) helper.ResponseWriter
}

func NewBookingHandler(bookingSvc service.BookingService) BookingHandler {
	return &Booking{
		bookingSvc: bookingSvc,
	}
}

func (b Booking) Book(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{Context: context}
	var bookingData models.Booking
	parseError := context.BindJSON(&bookingData)
	if parseError != nil {
		return w.Message(http.StatusBadRequest, "Invalid booking spec")
	}
	err := b.bookingSvc.Book(bookingData)
	if err != nil {
		w.Message(http.StatusInternalServerError, "There was an unexpected error during creating a booking")
	} else {
		w.Data(http.StatusOK, nil)
	}
	return w
}

func (b Booking) Fetch(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{Context: context}
	bookingId, exist := context.Params.Get("id")
	if !exist {
		return w.Message(http.StatusBadRequest, "Id not provided")
	}

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
	return w
}

func (b Booking) Modify(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{Context: context}
	bookingId, idExist := context.Params.Get("id")
	status, statusExist := context.Params.Get("status")
	if !(idExist || statusExist) {
		return w.Message(http.StatusBadRequest, "Either booking id or status is not provided")
	}
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
	return w
}

func (b Booking) Cancel(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{Context: context}
	bookingId, exist := context.Params.Get("id")
	reason, reasonExist := context.Params.Get("reason")
	if !(exist || reasonExist) {
		return w.Message(http.StatusBadRequest, "Either booking id or reason is not provided")
	}
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
	return w
}
