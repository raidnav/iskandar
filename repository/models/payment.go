package models

import "time"

type Payment struct {
	Id          int
	UserId      string
	AccountNo   string
	Amount      float32
	Method      string
	Status      string
	PaymentTime time.Time
	Notes       string
}
