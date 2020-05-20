package postgres

import (
	"errors"
	"github.com/code-and-chill/iskandar/constants"
	"github.com/code-and-chill/iskandar/repositories/models"
	"reflect"
)
import "github.com/jinzhu/gorm"

type booking struct {
	db *gorm.DB
}

func (t *booking) Fetch(id int) models.Booking {
	var booking = models.Booking{}
	t.db.Where("id = ?", id).First(&booking)
	return booking
}

func (t *booking) Create(bookingSpec models.Booking) error {
	return t.db.Create(&bookingSpec).Error //TODO: need to catch idempotent error separately
}

func (t *booking) Modify(id int, status string) error {
	booking := models.Booking{}
	t.db.Where("id = ?", id).First(&booking)
	if (reflect.DeepEqual(booking, models.Booking{})) {
		return errors.New(constants.DbNotFound)
	}
	booking.Status = status
	return t.db.Save(booking).Error
}

func NewBookingSchema(db *gorm.DB) *booking {
	db.AutoMigrate(&models.Booking{})
	return &booking{
		db: db,
	}
}
