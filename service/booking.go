package service

import (
	"errors"
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository"
	"github.com/code-and-chill/iskandar/repository/models"
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
	bookingAccessor repository.Booking
	logger          logrus.FieldLogger
}

func (b *Booking) Book(bookingSpec models.Booking) error {
	return b.bookingAccessor.Create(bookingSpec)
}

func (b *Booking) Fetch(id int) (models.Booking, error) {
	booking := b.bookingAccessor.Fetch(id)
	if reflect.DeepEqual(booking, models.Booking{}) {
		return models.Booking{}, errors.New(constant.DbNotFound)
	}
	return booking, nil
}

func (b *Booking) Modify(id int, status string) error {
	return b.bookingAccessor.Modify(id, status)
}

func (b *Booking) Cancel(id int, status string) error {
	return b.bookingAccessor.Modify(id, status)
}

func NewBookingService(bookingAccessor repository.Booking, logger logrus.FieldLogger) BookingService {
	return &Booking{
		bookingAccessor: bookingAccessor,
		logger:          logger,
	}
}
