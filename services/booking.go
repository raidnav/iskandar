package services

import (
	"errors"
	"github.com/code-and-chill/iskandar/constants"
	"github.com/code-and-chill/iskandar/repositories"
	"github.com/code-and-chill/iskandar/repositories/models"
	"github.com/sirupsen/logrus"
	"reflect"
)

type BookingService interface {
	Book(bookingSpec models.Booking) error
	Fetch(id int) (models.Booking, error)
	Modify(id int, status string) error
	Cancel(id int, reason string) error
}

type Booking struct {
	bookingAccessor repositories.Booking
	logger          logrus.FieldLogger
}

func (b *Booking) Book(bookingSpec models.Booking) error {

	return b.bookingAccessor.Create(bookingSpec)
}

func (b *Booking) Fetch(id int) (models.Booking, error) {
	booking := b.bookingAccessor.Fetch(id)
	if reflect.DeepEqual(booking, models.Booking{}) {
		return models.Booking{}, errors.New(constants.DbNotFound)
	}
	return booking, nil
}

func (b *Booking) Modify(id int, status string) error {
	return b.bookingAccessor.Modify(id, status)
}

func (b *Booking) Cancel(id int, status string) error {
	return b.bookingAccessor.Modify(id, status)
}

func NewBookingService(bookingAccessor repositories.Booking, logger logrus.FieldLogger) BookingService {
	return &Booking{
		bookingAccessor: bookingAccessor,
		logger:          logger,
	}
}
