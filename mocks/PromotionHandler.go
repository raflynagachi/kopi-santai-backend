// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// PromotionHandler is an autogenerated mock type for the PromotionHandler type
type PromotionHandler struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: c
func (_m *PromotionHandler) FindAll(c *gin.Context) {
	_m.Called(c)
}

// FindAllUnscoped provides a mock function with given fields: c
func (_m *PromotionHandler) FindAllUnscoped(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewPromotionHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromotionHandler creates a new instance of PromotionHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromotionHandler(t mockConstructorTestingTNewPromotionHandler) *PromotionHandler {
	mock := &PromotionHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
