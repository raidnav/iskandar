package service

import (
	"errors"
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/sirupsen/logrus"
	"reflect"
)

type TicketService interface {
	Create(TicketSpec models.Ticket) error
	Fetch(id int) (models.Ticket, error)
}

type Ticket struct {
	TicketAccessor repository.Ticket
	logger         logrus.FieldLogger
}

func (b *Ticket) Create(TicketSpec models.Ticket) error {
	return b.TicketAccessor.Create(TicketSpec)
}

func (b *Ticket) Fetch(id int) (models.Ticket, error) {
	Ticket := b.TicketAccessor.Fetch(id)
	if reflect.DeepEqual(Ticket, models.Ticket{}) {
		return models.Ticket{}, errors.New(constant.DbNotFound)
	}
	return Ticket, nil
}

func NewTicketService(TicketAccessor repository.Ticket, logger logrus.FieldLogger) TicketService {
	return &Ticket{
		TicketAccessor: TicketAccessor,
		logger:         logger,
	}
}
