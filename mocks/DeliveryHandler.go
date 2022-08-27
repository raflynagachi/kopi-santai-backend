// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// DeliveryHandler is an autogenerated mock type for the DeliveryHandler type
type DeliveryHandler struct {
	mock.Mock
}

// UpdateStatus provides a mock function with given fields: c
func (_m *DeliveryHandler) UpdateStatus(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewDeliveryHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeliveryHandler creates a new instance of DeliveryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeliveryHandler(t mockConstructorTestingTNewDeliveryHandler) *DeliveryHandler {
	mock := &DeliveryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}