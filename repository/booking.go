package repository

import "github.com/code-and-chill/iskandar/repository/models"

type Booking interface {
	Fetch(id int) models.Booking
	Create(bookingSpec models.Booking) error
	Modify(id int, status string) error // only status allowed to be modified
}
