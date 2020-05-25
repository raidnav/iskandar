package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/handler/helper"
	"github.com/code-and-chill/iskandar/lib"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Ticket struct {
	bookingSvc service.BookingService
	ticketSvc  service.TicketService
}

type TicketHandler interface {
	Save(context *gin.Context) *gin.Context
	Fetch(context *gin.Context) *gin.Context
}

func NewTicketHandler(ticketSvc service.TicketService, bookingSvc service.BookingService) TicketHandler {
	return &Ticket{
		ticketSvc:  ticketSvc,
		bookingSvc: bookingSvc,
	}
}

func (t Ticket) Save(context *gin.Context) *gin.Context {
	w := helper.ResponseWriter{Context: context}
	id, exist := context.Params.Get("booking_id")
	if !exist {
		return w.Message(http.StatusBadRequest, "id not provided")
	}

	bookingId, parseError := strconv.Atoi(id)
	if parseError != nil {
		return w.Message(http.StatusNotFound, "bookingSvc id can't be parsed to number.")
	}

	_, err := t.bookingSvc.Fetch(bookingId)
	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			w.Message(http.StatusMethodNotAllowed, "you can't proceed without giving correct booking_id")
		default:
			w.Message(http.StatusInternalServerError, "there's error while fetching bookingSvc data")
		}
		return w.Context
	}

	lib.GeneratePdf()

	ticket := models.Ticket{
		Id:       bookingId,
		FilePath: "",
	}
	err = t.ticketSvc.Create(ticket)
	if err != nil {
		w.Message(http.StatusInternalServerError, "")
	}
	return w.Data(http.StatusOK, nil)
}

func (t Ticket) Fetch(context *gin.Context) *gin.Context {
	w := helper.ResponseWriter{Context: context}
	documentId, exist := context.Params.Get("documentId")
	if !exist {
		return w.Message(http.StatusBadRequest, "Id not provided")
	}

	id, numError := strconv.Atoi(documentId)
	if numError != nil {
		return w.Message(http.StatusBadRequest, "Unable to convert documentId: "+documentId)
	}

	ticket, err := t.ticketSvc.Fetch(id)

	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			w.Message(http.StatusNotFound, "ticketSvc Id was not found")
		default:
			w.Message(http.StatusInternalServerError, "There was an unexpected error during getting ticketSvc data")
		}
		return w.Context
	}
	return w.Data(http.StatusOK, ticket)
}
