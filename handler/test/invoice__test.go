package test

import (
	"github.com/code-and-chill/iskandar/handler"
	mockservice "github.com/code-and-chill/iskandar/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestInvoiceHandler_Save(t *testing.T) {
	h := handler.NewInvoiceHandler(mockservice.NewMockInvoiceService(gomock.NewController(t)))
	ctx := gin.Context{}
	h.Save(&ctx)
}

func TestInvoiceHandler_Fetch(t *testing.T) {
	h := handler.NewInvoiceHandler(mockservice.NewMockInvoiceService(gomock.NewController(t)))
	ctx := gin.Context{}
	h.Fetch(&ctx)
}
