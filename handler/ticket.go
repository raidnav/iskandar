package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/helper"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TicketHandler struct {
	bookingSvc service.BookingService
	ticketSvc  service.TicketService
}

type TicketController interface {
	Save() gin.HandlerFunc
	Fetch() gin.HandlerFunc
}

func NewTicketHandler(ticketSvc service.TicketService, bookingSvc service.BookingService) TicketController {
	return &TicketHandler{
		ticketSvc:  ticketSvc,
		bookingSvc: bookingSvc,
	}
}

func (t TicketHandler) Save() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		id, exist := context.Params.Get("booking_id")
		if !exist {
			w.Message(http.StatusBadRequest, "id not provided")
		}

		bookingId, parseError := strconv.Atoi(id)
		if parseError != nil {
			w.Message(http.StatusNotFound, "booking id can't be parsed to number.")
			return
		}

		_, err := t.bookingSvc.Fetch(bookingId)
		if err != nil {
			switch err.Error() {
			case constant.DbNotFound:
				w.Message(http.StatusMethodNotAllowed, "you can't proceed without giving correct booking_id")
			default:
				w.Message(http.StatusInternalServerError, "there's error while fetching booking data")
			}
			return
		}

		helper.GeneratePdf()

		ticket := models.Ticket{
			Id:       bookingId,
			FilePath: "",
		}
		err = t.ticketSvc.Create(ticket)
		if err != nil {
			w.Message(http.StatusInternalServerError, "")
		}
	}
}

func (t TicketHandler) Fetch() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		documentId, exist := context.Params.Get("documentId")
		if !exist {
			w.Message(http.StatusBadRequest, "Id not provided")
			return
		}

		id, numError := strconv.Atoi(documentId)
		if numError != nil {
			w.Message(http.StatusBadRequest, "Unable to convert documentId: "+documentId)
			return
		}

		ticket, err := t.ticketSvc.Fetch(id)

		if err != nil {
			switch err.Error() {
			case constant.DbNotFound:
				w.Message(http.StatusNotFound, "ticket Id was not found")
			default:
				w.Message(http.StatusInternalServerError, "There was an unexpected error during getting ticket data")
			}
			return
		}
		w.Data(http.StatusOK, ticket)
	}
}
