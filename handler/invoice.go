package handler

import (
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceSvc service.InvoiceService
}

type InvoiceController interface {
	Save() gin.HandlerFunc
	Fetch() gin.HandlerFunc
}

func NewInvoiceHandler(invoiceSvc service.InvoiceService) InvoiceController {
	return &InvoiceHandler{
		invoiceSvc: invoiceSvc,
	}
}

func (i InvoiceHandler) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: impl
	}
}

func (i InvoiceHandler) Fetch() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO: impl
	}
}
