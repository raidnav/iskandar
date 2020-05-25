package test

import (
	"encoding/json"
	"github.com/code-and-chill/iskandar/handler"
	"github.com/code-and-chill/iskandar/repository/models"
	mockService "github.com/code-and-chill/iskandar/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestBookingHandler_Book(t *testing.T) {
	h := handler.NewBookingHandler(mockService.NewMockBookingService(gomock.NewController(t)))
	payload, _ := json.Marshal(models.Booking{
		Id:     1,
		UserId: "Dev",
		Detail: []models.BookingDetail{
			{
				"",
				models.BookingFare{},
				[]models.BookingPassenger{
					{
						"",
						"",
						"",
					},
				},
			},
			{
				"",
				models.BookingFare{},
				[]models.BookingPassenger{
					{
						"",
						"",
						"",
					},
				},
			},
		},
		TotalFare: 5500.0,
		Status:    "BOOKED",
		Notes:     "",
	})
	ctx := gin.Context{
		Request: &http.Request{
			Method: http.MethodPost,
			Header: http.Header{},
			Body:   nil,
		},
		Params: gin.Params{
			gin.Param{
				Key:   "booking",
				Value: string(payload),
			},
		},
	}
	h.Book(&ctx)
}

func TestBookingHandler_Fetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mockService.NewMockBookingService(ctrl)
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
	bookSvc := mockService.NewMockBookingService(ctrl)
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

func TestGivenExistedBookingIdShouldAbleToCancel(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mockService.NewMockBookingService(ctrl)
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

func TestGivenNonExistedBookingIdShouldFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookSvc := mockService.NewMockBookingService(ctrl)
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
