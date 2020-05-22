package handler

import (
	"github.com/code-and-chill/iskandar/constant"
	"github.com/code-and-chill/iskandar/repository/models"
	"github.com/code-and-chill/iskandar/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PaymentHandler struct {
	PaymentSvc service.PaymentService
}

type PaymentController interface {
	GenerateRequestSpec() gin.HandlerFunc
	Pay() gin.HandlerFunc
	Cancel() gin.HandlerFunc
}

func NewPaymentHandler(PaymentSvc service.PaymentService) PaymentController {
	return &PaymentHandler{
		PaymentSvc: PaymentSvc,
	}
}

func (b PaymentHandler) GenerateRequestSpec() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
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
				"2. Go to payment\n" +
				"3. Select other providers\n4. Add above numbers.\n" +
				"5. Pay",
		}
		w.Data(http.StatusOK, resp)
	}
}

func (b PaymentHandler) Pay() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
		var paymentData models.Payment
		parseError := context.BindJSON(&paymentData)
		if parseError != nil {
			w.Message(http.StatusBadRequest, "Invalid Payment spec")
		}
		err := b.PaymentSvc.Pay(paymentData)
		if err != nil {
			if err.Error() == constant.UnsuccessfulPayment {
				w.Message(http.StatusNotModified, "Client hasn't complete the payment")
				return
			}
			w.Message(http.StatusInternalServerError, "There was an unexpected error during creating a Payment")
			return
		}
		w.Data(http.StatusOK, nil)
	}
}

func (b PaymentHandler) Cancel() gin.HandlerFunc {
	return func(context *gin.Context) {
		w := ResponseWriter{ctx: context}
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

		err := b.PaymentSvc.Cancel(id)
		if err != nil {
			switch err.Error() {
			case constant.DbNotFound:
				w.Message(http.StatusNotFound, "Payment id is not exist in our data store")
			default:
				w.Message(http.StatusInternalServerError, "There's an error when updating payment id: "+paymentId)
			}
			return
		}
		w.Data(http.StatusOK, nil)
	}
}
