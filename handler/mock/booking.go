// Code generated by MockGen. DO NOT EDIT.
// Source: handler/booking.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBookingController is a mock of BookingController interface
type MockBookingController struct {
	ctrl     *gomock.Controller
	recorder *MockBookingControllerMockRecorder
}

// MockBookingControllerMockRecorder is the mock recorder for MockBookingController
type MockBookingControllerMockRecorder struct {
	mock *MockBookingController
}

// NewMockBookingController creates a new mock instance
func NewMockBookingController(ctrl *gomock.Controller) *MockBookingController {
	mock := &MockBookingController{ctrl: ctrl}
	mock.recorder = &MockBookingControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookingController) EXPECT() *MockBookingControllerMockRecorder {
	return m.recorder
}

// Book mocks base method
func (m *MockBookingController) Book() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Book")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Book indicates an expected call of Book
func (mr *MockBookingControllerMockRecorder) Book() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockBookingController)(nil).Book))
}

// Fetch mocks base method
func (m *MockBookingController) Fetch() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (mr *MockBookingControllerMockRecorder) Fetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockBookingController)(nil).Fetch))
}

// Modify mocks base method
func (m *MockBookingController) Modify() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Modify indicates an expected call of Modify
func (mr *MockBookingControllerMockRecorder) Modify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockBookingController)(nil).Modify))
}

// Cancel mocks base method
func (m *MockBookingController) Cancel() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockBookingControllerMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockBookingController)(nil).Cancel))
}