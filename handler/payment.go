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
	GenerateRequestSpec(context *gin.Context)
	Pay(context *gin.Context)
	Cancel(context *gin.Context)
}

func NewPaymentHandler(paymentSvc service.PaymentService) PaymentHandler {
	return &Payment{
		paymentSvc: paymentSvc,
	}
}

func (p Payment) GenerateRequestSpec(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	userAccountNo, keyFound := context.Params.Get("accountNo")
	if keyFound != false {
		w.Message(http.StatusBadRequest, "accountNo is not found")
		return
	}
	resp := struct {
		account    string
		stepByStep string
	}{
		account: "311" + userAccountNo,
		stepByStep: "1. Open up your mobile banking app\n" +
			"2. Go to paymentSvc\n" +
			"3. Select other providers\n4. Add above numbers.\n" +
			"5. Pay",
	}
	w.Data(http.StatusOK, resp)
}

func (p Payment) Pay(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	var paymentData models.Payment
	parseError := context.BindJSON(&paymentData)
	if parseError != nil {
		w.Message(http.StatusBadRequest, "Invalid Payment spec")
		return
	}
	err := p.paymentSvc.Pay(paymentData)
	if err != nil {
		switch err.Error() {
		case constant.UnsuccessfulPayment:
			w.Message(http.StatusNotModified, "Client hasn't complete the paymentSvc")
		default:
			w.Message(http.StatusInternalServerError, "There was an unexpected error during creating a Payment")
		}
		return
	}
	w.Data(http.StatusOK, nil)
}

func (p Payment) Cancel(context *gin.Context) {
	w := helper.ResponseWriter{Context: context}
	paymentId, keyFound := context.Params.Get("id")
	if !keyFound {
		w.Message(http.StatusBadRequest, "Payment id is not provided")
		return
	}
	id, parseErr := strconv.Atoi(paymentId)
	if parseErr != nil {
		w.Message(http.StatusBadRequest, "Payment seems contains ineligible character.")
		return
	}

	err := p.paymentSvc.Cancel(id)
	if err != nil {
		switch err.Error() {
		case constant.DbNotFound:
			w.Message(http.StatusNotFound, "Payment id is not exist in our data store")
		default:
			w.Message(http.StatusInternalServerError, "There's an error when updating paymentSvc id: "+paymentId)
		}
		return
	}
	w.Data(http.StatusOK, nil)
}
