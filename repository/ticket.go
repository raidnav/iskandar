package repository

import "github.com/code-and-chill/iskandar/repository/models"

type Ticket interface {
	Fetch(id int) models.Ticket
	Create(TicketSpec models.Ticket) error
}
