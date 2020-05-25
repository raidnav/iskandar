package test

import (
	"github.com/code-and-chill/iskandar/handler"
	mockservice "github.com/code-and-chill/iskandar/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPaymentHandler_GenerateRequestSpec(t *testing.T) {
	h := handler.NewPaymentHandler(mockservice.NewMockPaymentService(gomock.NewController(t)))
	ctx := gin.Context{}
	h.GenerateRequestSpec(&ctx)
}

func TestPaymentHandler_Pay(t *testing.T) {
	h := handler.NewPaymentHandler(mockservice.NewMockPaymentService(gomock.NewController(t)))
	ctx := gin.Context{}
	h.Pay(&ctx)
}

func TestPaymentHandler_Cancel(t *testing.T) {
	h := handler.NewPaymentHandler(mockservice.NewMockPaymentService(gomock.NewController(t)))
	ctx := gin.Context{}
	h.Cancel(&ctx)
}
