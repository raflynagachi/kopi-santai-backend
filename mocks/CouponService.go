// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// CouponService is an autogenerated mock type for the CouponService type
type CouponService struct {
	mock.Mock
}

// Create provides a mock function with given fields: req
func (_m *CouponService) Create(req *dto.CouponPostReq) (*dto.CouponRes, error) {
	ret := _m.Called(req)

	var r0 *dto.CouponRes
	if rf, ok := ret.Get(0).(func(*dto.CouponPostReq) *dto.CouponRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CouponRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.CouponPostReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: id
func (_m *CouponService) DeleteByID(id uint) (gin.H, error) {
	ret := _m.Called(id)

	var r0 gin.H
	if rf, ok := ret.Get(0).(func(uint) gin.H); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.H)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *CouponService) FindAll() ([]*dto.CouponRes, error) {
	ret := _m.Called()

	var r0 []*dto.CouponRes
	if rf, ok := ret.Get(0).(func() []*dto.CouponRes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.CouponRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCouponByUserID provides a mock function with given fields: userID
func (_m *CouponService) FindCouponByUserID(userID uint) ([]*dto.CouponRes, error) {
	ret := _m.Called(userID)

	var r0 []*dto.CouponRes
	if rf, ok := ret.Get(0).(func(uint) []*dto.CouponRes); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.CouponRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCouponService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCouponService creates a new instance of CouponService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCouponService(t mockConstructorTestingTNewCouponService) *CouponService {
	mock := &CouponService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
