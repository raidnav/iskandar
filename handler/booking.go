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
	Book(context *gin.Context)
	Fetch(context *gin.Context)
	Modify(context *gin.Context)
	Cancel(context *gin.Context)
}

func NewBookingHandler(bookingSvc service.BookingService) BookingHandler {
	return &Booking{
		bookingSvc: bookingSvc,
	}
}

func (b Booking) Book(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	var bookingData models.Booking
	parseError := context.BindJSON(&bookingData)
	if parseError != nil {
		w.Message(http.StatusBadRequest, "Invalid bookingSvc spec. \n"+parseError.Error())
		return
	}
	err := b.bookingSvc.Book(bookingData)
	if err != nil {
		switch err.Error() {
		case "pq: duplicate key value violates unique constraint \"bookings_pkey\"":
			w.Message(http.StatusUnprocessableEntity, "Duplicate data requested")
		default:
			w.Message(http.StatusInternalServerError, err.Error())
		}
	} else {
		w.Data(http.StatusOK, nil)
	}
}

func (b Booking) Fetch(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	bookingId, exist := context.Params.Get("id")
	if !exist {
		w.Message(http.StatusBadRequest, "Id not provided")
		return
	}

	id, numError := strconv.Atoi(bookingId)
	if numError != nil {
		w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
		return
	}

	booking, err := b.bookingSvc.Fetch(id)

	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			w.Message(http.StatusNotFound, "Booking Id was not found")
		default:
			w.Message(http.StatusInternalServerError, "There was an unexpected error during getting bookingSvc data")
		}
	} else {
		w.Data(http.StatusOK, booking)
	}
}

func (b Booking) Modify(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	bookingId, idExist := context.Params.Get("id")
	status := context.PostForm("status")
	if !(idExist || status != "") {
		w.Message(http.StatusBadRequest, "Either bookingSvc id or status is not provided")
		return
	}
	id, numError := strconv.Atoi(bookingId)
	if numError != nil {
		w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
	}

	err := b.bookingSvc.Modify(id, status)
	if err != nil {
		w.Message(http.StatusInternalServerError, "There was an unexpected error during getting bookingSvc data")
	} else {
		w.Data(http.StatusOK, nil)
	}
}

func (b Booking) Cancel(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	bookingId, exist := context.Params.Get("id")
	reason := context.PostForm("reason")
	if !(exist || reason != "") {
		w.Message(http.StatusBadRequest, "Either bookingSvc id or reason is not provided")
		return
	}
	id, numError := strconv.Atoi(bookingId)
	if numError != nil {
		w.Message(http.StatusBadRequest, "Unable to convert bookingId: "+bookingId)
		return
	}

	err := b.bookingSvc.Cancel(id, reason)

	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			w.Message(http.StatusNotFound, "Requested bookingSvc id was not found")
		default:
			w.Message(http.StatusInternalServerError, "There was an unexpected error during cancelling bookingSvc")
		}
	} else {
		w.Data(http.StatusNoContent, nil)
	}
}
