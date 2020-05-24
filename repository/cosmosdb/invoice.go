package cosmosdb

import (
	"github.com/code-and-chill/iskandar/repository/models"
	"gopkg.in/mgo.v2"
)

type invoice struct {
	table *mgo.Session
}

func (i invoice) Save(invoice models.Invoice) error {
	panic("implement me")
}

func (i invoice) Fetch() models.Invoice {
	panic("implement me")
}

func NewInvoiceCollection(mongo *mgo.Session) *invoice {
	return &invoice{
		table: mongo,
	}
}
