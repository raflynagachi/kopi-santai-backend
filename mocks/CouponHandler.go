// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// CouponHandler is an autogenerated mock type for the CouponHandler type
type CouponHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *CouponHandler) Create(c *gin.Context) {
	_m.Called(c)
}

// DeleteByID provides a mock function with given fields: c
func (_m *CouponHandler) DeleteByID(c *gin.Context) {
	_m.Called(c)
}

// FindAll provides a mock function with given fields: c
func (_m *CouponHandler) FindAll(c *gin.Context) {
	_m.Called(c)
}

// FindCouponByUserID provides a mock function with given fields: c
func (_m *CouponHandler) FindCouponByUserID(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewCouponHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewCouponHandler creates a new instance of CouponHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCouponHandler(t mockConstructorTestingTNewCouponHandler) *CouponHandler {
	mock := &CouponHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
