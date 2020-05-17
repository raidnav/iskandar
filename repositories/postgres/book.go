package postgres

import (
	"github.com/code-and-chill/iskandar/repositories"
	"github.com/code-and-chill/iskandar/repositories/models"
	"github.com/sirupsen/logrus"
)
import "github.com/jinzhu/gorm"

type booking struct {
	dbAdapter *gorm.DB
	log       logrus.FieldLogger
}

func (b *booking) Fetch(id int) (models.Booking, error) {
	panic("implement me")
}

func (b *booking) Create(bookingSpec models.Booking) error {
	panic("implement me")
}

func (b *booking) Modify(id int, status string) error {
	panic("implement me")
}

func (b *booking) Cancel(id int) error {
	panic("implement me")
}

func NewPostgresBookingSchema(db *gorm.DB, log logrus.FieldLogger) repositories.Booking {
	return &booking{
		dbAdapter: db,
		log:       log,
	}
}
