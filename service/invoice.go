package service

import (
	"github.com/code-and-chill/iskandar/repository"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/sirupsen/logrus"
)

type InvoiceService interface {
	Create(invoiceSpec models.Invoice) error
	Fetch(id int) (models.Invoice, error)
}

type Invoice struct {
	InvoiceAccessor repository.Invoice
	logger          logrus.FieldLogger
}

func (i Invoice) Create(invoiceSpec models.Invoice) error {
	panic("implement me")
}

func (i Invoice) Fetch(id int) (models.Invoice, error) {
	panic("implement me")
}

func NewInvoiceService(invoiceAccessor repository.Invoice, logger logrus.FieldLogger) InvoiceService {
	return &Invoice{
		InvoiceAccessor: invoiceAccessor,
		logger:          logger,
	}
}
