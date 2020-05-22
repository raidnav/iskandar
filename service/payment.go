package service

import (
	"errors"
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/helper"
	pg "github.com/code-and-chill/iskandar/provider/payment-gateway"
	"github.com/code-and-chill/iskandar/repository"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/sirupsen/logrus"
	"time"
)

type PaymentService interface {
	Pay(PaymentSpec models.Payment) error
	Cancel(id int) error
	processTransfer()
	assignIncomingPayment()
}

type Payment struct {
	paymentAccessor repository.Payment
	logger          logrus.FieldLogger
}

func (p Payment) Pay(paymentSpec models.Payment) error {
	switch paymentSpec.Method {
	case constant.BankTransfer:
		panic("Implement me")
	case constant.VirtualAccount:
		resp := pg.Request(paymentSpec.AccountNo, paymentSpec.Amount)
		if resp.Success != true {
			return errors.New(constant.UnsuccessfulPayment)
		}
	default:
		panic("Implement me")
	}
	return p.paymentAccessor.Create(paymentSpec)
}

func (p Payment) Cancel(id int) error {
	return p.paymentAccessor.Cancel(id)
}

func (p Payment) processTransfer() {
	panic("implement me")
}

func (p Payment) assignIncomingPayment() {
	from := helper.ToMillis(time.Now().Add(-3 * 24 * time.Hour))
	to := helper.ToMillis(time.Now())

	p.paymentAccessor.Fetch(from, to)
	// TODO continued later
}

func NewPaymentService(PaymentAccessor repository.Payment, logger logrus.FieldLogger) PaymentService {
	return &Payment{
		paymentAccessor: PaymentAccessor,
		logger:          logger,
	}
}
