package postgres

import (
	"errors"
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository/models"
	"reflect"
)
import "github.com/jinzhu/gorm"

type Booking struct {
	db *gorm.DB
}

func (t *Booking) Fetch(id int) models.Booking {
	var booking = models.Booking{}
	t.db.Where("id = ?", id).First(&booking)
	return booking
}

func (t *Booking) Create(bookingSpec models.Booking) error {
	return t.db.Create(&bookingSpec).Error //TODO: need to catch idempotent error separately
}

func (t *Booking) Modify(id int, status string) error {
	booking := models.Booking{}
	t.db.Where("id = ?", id).First(&booking)
	if (reflect.DeepEqual(booking, models.Booking{})) {
		return errors.New(constant.DbNotFound)
	}
	booking.Status = status
	return t.db.Save(booking).Error
}

func NewBookingSchema(db *gorm.DB) *Booking {
	db.AutoMigrate(&models.Booking{})
	return &Booking{
		db: db,
	}
}
