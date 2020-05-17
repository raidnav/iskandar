package repositories

import "github.com/code-and-chill/iskandar/repositories/models"

type Booking interface {
	Fetch(id int) (models.Booking, error)
	Create(bookingSpec models.Booking) error
	Modify(id int, status string) error // only status allowed to be modified
	Cancel(id int) error
}
