package postgres

import (
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/jinzhu/gorm"
)

type Invoice struct {
	db *gorm.DB
}

func (i *Invoice) Save(invoice models.Invoice) error {
	panic("Implement me")
}

func (i *Invoice) Fetch() models.Invoice {
	panic("Implement me")
}

func NewInvoiceSchema(db *gorm.DB) *Invoice {
	db.AutoMigrate(&models.Invoice{})
	return &Invoice{
		db: db,
	}
}
