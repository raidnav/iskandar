package services

import (
	"fmt"
	"github.com/code-and-chill/iskandar/repositories"
	"github.com/code-and-chill/iskandar/repositories/models"
	"github.com/sirupsen/logrus"
)

type BookingService interface {
	Book(bookingSpec models.Booking) error
	Fetch(id int) (models.Booking, error)
	Modify(id int, status string) (models.Booking, error)
	Cancel(id int, reason string) (models.Booking, error)
}

type Booking struct {
	accessor repositories.Booking
	logger   logrus.FieldLogger
}

func (b *Booking) Book(bookingSpec models.Booking) error {
	// TODO: implement real functionality
	logrus.Info(fmt.Sprintf("curret parser: %d-%s-%s",
		bookingSpec.Id,
		bookingSpec.Status,
		bookingSpec.UserId))
	return nil
}

func (b *Booking) Fetch(id int) (models.Booking, error) {
	panic("implement me")
}

func (b *Booking) Modify(id int, status string) (models.Booking, error) {
	panic("implement me")
}

func (b *Booking) Cancel(id int, reason string) (models.Booking, error) {
	panic("implement me")
}

func NewBookingService(bookingAccessor repositories.Booking, logger logrus.FieldLogger) BookingService {
	return &Booking{
		accessor: bookingAccessor,
		logger:   logger,
	}
}
