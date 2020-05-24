package test

import (
	"github.com/code-and-chill/iskandar/handler"
	mock_service "github.com/code-and-chill/iskandar/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestBookingHandler_Book(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mock_service.NewMockBookingService(ctrl)
	h := handler.NewBookingHandler(bookSvc)
	ctx := gin.Context{
		Params: gin.Params{
			gin.Param{
				Key:   "",
				Value: "",
			},
		},
	}
	h.Book(&ctx)
}

func TestBookingHandler_Fetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mock_service.NewMockBookingService(ctrl)
	h := handler.NewBookingHandler(bookSvc)
	ctx := gin.Context{
		Params: gin.Params{
			gin.Param{
				Key:   "id",
				Value: "1",
			},
		},
	}
	h.Fetch(&ctx)
}

func TestBookingHandler_Modify(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mock_service.NewMockBookingService(ctrl)
	h := handler.NewBookingHandler(bookSvc)
	ctx := gin.Context{
		Params: gin.Params{
			gin.Param{
				Key:   "id",
				Value: "1",
			},
		},
	}
	h.Modify(&ctx)
}

func TestBookingHandler_Cancel(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mock_service.NewMockBookingService(ctrl)
	h := handler.NewBookingHandler(bookSvc)
	ctx := gin.Context{
		Params: gin.Params{
			gin.Param{
				Key:   "id",
				Value: "1",
			},
		},
	}
	h.Cancel(&ctx)

}
