// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// OrderHandler is an autogenerated mock type for the OrderHandler type
type OrderHandler struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: c
func (_m *OrderHandler) CreateOrder(c *gin.Context) {
	_m.Called(c)
}

// FindAll provides a mock function with given fields: c
func (_m *OrderHandler) FindAll(c *gin.Context) {
	_m.Called(c)
}

// FindByUserID provides a mock function with given fields: c
func (_m *OrderHandler) FindByUserID(c *gin.Context) {
	_m.Called(c)
}

// FindOrderByIDAndUserID provides a mock function with given fields: c
func (_m *OrderHandler) FindOrderByIDAndUserID(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewOrderHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderHandler creates a new instance of OrderHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderHandler(t mockConstructorTestingTNewOrderHandler) *OrderHandler {
	mock := &OrderHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}