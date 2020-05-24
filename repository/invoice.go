package repository

import "github.com/code-and-chill/iskandar/repository/models"

type Invoice interface {
	Save(invoice models.Invoice) error
	Fetch() models.Invoice
}
