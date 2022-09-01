// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// PromotionService is an autogenerated mock type for the PromotionService type
type PromotionService struct {
	mock.Mock
}

// CreatePromotion provides a mock function with given fields: req
func (_m *PromotionService) CreatePromotion(req *dto.PromotionPostReq) (*dto.PromotionRes, error) {
	ret := _m.Called(req)

	var r0 *dto.PromotionRes
	if rf, ok := ret.Get(0).(func(*dto.PromotionPostReq) *dto.PromotionRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PromotionRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.PromotionPostReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePromotionByID provides a mock function with given fields: id
func (_m *PromotionService) DeletePromotionByID(id uint) (gin.H, error) {
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
func (_m *PromotionService) FindAll() ([]*dto.PromotionRes, error) {
	ret := _m.Called()

	var r0 []*dto.PromotionRes
	if rf, ok := ret.Get(0).(func() []*dto.PromotionRes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.PromotionRes)
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

// FindAllUnscoped provides a mock function with given fields:
func (_m *PromotionService) FindAllUnscoped() ([]*dto.PromotionRes, error) {
	ret := _m.Called()

	var r0 []*dto.PromotionRes
	if rf, ok := ret.Get(0).(func() []*dto.PromotionRes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.PromotionRes)
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

type mockConstructorTestingTNewPromotionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromotionService creates a new instance of PromotionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromotionService(t mockConstructorTestingTNewPromotionService) *PromotionService {
	mock := &PromotionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
