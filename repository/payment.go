package repository

import "github.com/code-and-chill/iskandar/repository/models"

type Payment interface {
	FetchId(id int) models.Payment
	Fetch(from int64, to int64) []models.Payment
	Create(PaymentSpec models.Payment) error
	Cancel(id int) error
}
