package postgres

import (
	"errors"
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository/models"
	"reflect"
)
import "github.com/jinzhu/gorm"

type Payment struct {
	db *gorm.DB
}

func (t *Payment) FetchId(id int) models.Payment {
	var Payment = models.Payment{}
	t.db.Where("id = ?", id).First(&Payment)
	return Payment
}

func (t *Payment) Cancel(id int) error {
	Payment := models.Payment{}
	t.db.Where("id = ?", id).First(&Payment)
	if (reflect.DeepEqual(Payment, models.Payment{})) {
		return errors.New(constant.DbNotFound)
	}
	return t.db.Save(Payment).Error
}

func (t *Payment) Fetch(from int64, to int64) []models.Payment {
	var Payment []models.Payment
	t.db.Where("from >= ? and to <= to", from, to).First(&Payment)
	return Payment
}

func (t *Payment) Create(PaymentSpec models.Payment) error {
	return t.db.Create(&PaymentSpec).Error //TODO: need to catch idempotent error separately
}

func NewPaymentSchema(db *gorm.DB) *Payment {
	db.AutoMigrate(&models.Payment{})
	return &Payment{
		db: db,
	}
}
