package postgres

import (
	"github.com/code-and-chill/iskandar/repository/models"
)
import "github.com/jinzhu/gorm"

type Ticket struct {
	db *gorm.DB
}

func (t *Ticket) Fetch(id int) models.Ticket {
	ticket := models.Ticket{}
	t.db.Where("id = ?", id).First(&ticket)
	return ticket
}

func (t *Ticket) Create(TicketSpec models.Ticket) error {
	return t.db.Create(&TicketSpec).Error //TODO: need to catch idempotent error separately
}

func NewTicketSchema(db *gorm.DB) *Ticket {
	db.AutoMigrate(&models.Ticket{})
	return &Ticket{
		db: db,
	}
}
