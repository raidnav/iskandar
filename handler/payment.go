package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/handler/helper"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Payment struct {
	paymentSvc service.PaymentService
}

type PaymentHandler interface {
	GenerateRequestSpec(context *gin.Context) helper.ResponseWriter
	Pay(context *gin.Context) helper.ResponseWriter
	Cancel(context *gin.Context) helper.ResponseWriter
}

func NewPaymentHandler(PaymentSvc service.PaymentService) PaymentHandler {
	return &Payment{
		paymentSvc: PaymentSvc,
	}
}

func (b Payment) GenerateRequestSpec(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{ctx: context}
	userAccountNo, keyFound := context.Params.Get("accountNo")
	if keyFound != false {
		return w.Message(http.StatusBadRequest, "accountNo is not found")
	}
	resp := struct {
		account    string
		stepByStep string
	}{
		account: "311" + userAccountNo,
		stepByStep: "1. Open up your mobile banking app\n" +
			"2. Go to payment\n" +
			"3. Select other providers\n4. Add above numbers.\n" +
			"5. Pay",
	}
	return w.Data(http.StatusOK, resp)
}

func (b Payment) Pay(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{ctx: context}
	var paymentData models.Payment
	parseError := context.BindJSON(&paymentData)
	if parseError != nil {
		w.Message(http.StatusBadRequest, "Invalid Payment spec")
	}
	err := b.paymentSvc.Pay(paymentData)
	if err != nil {
		if err.Error() == constant.UnsuccessfulPayment {
			w.Message(http.StatusNotModified, "Client hasn't complete the payment")
			return w
		}
		w.Message(http.StatusInternalServerError, "There was an unexpected error during creating a Payment")
		return w
	}
	w.Data(http.StatusOK, nil)
	return w
}

func (b Payment) Cancel(context *gin.Context) helper.ResponseWriter {
	w := helper.ResponseWriter{ctx: context}
	paymentId, keyFound := context.Params.Get("id")
	if !keyFound {
		return w.Message(http.StatusBadRequest, "Payment id is not provided")
	}
	id, parseErr := strconv.Atoi(paymentId)
	if parseErr != nil {
		return w.Message(http.StatusBadRequest, "Payment seems contains ineligible character.")
	}

	err := b.paymentSvc.Cancel(id)
	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			return w.Message(http.StatusNotFound, "Payment id is not exist in our data store")
		default:
			return w.Message(http.StatusInternalServerError, "There's an error when updating payment id: "+paymentId)
		}
	}
	return w.Data(http.StatusOK, nil)
}
