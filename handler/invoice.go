package handler

import (
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
)

type Invoice struct {
	invoiceSvc service.InvoiceService
}

func (i Invoice) Save(context *gin.Context) {
	panic("implement me")
}

func (i Invoice) Fetch(context *gin.Context) {
	panic("implement me")
}

type InvoiceController interface {
	Save(context *gin.Context)
	Fetch(context *gin.Context)
}

func NewInvoiceHandler(invoiceSvc service.InvoiceService) InvoiceController {
	return &Invoice{
		invoiceSvc: invoiceSvc,
	}
}
