// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/raflynagachi/kopi-santai-backend/dto"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	model "github.com/raflynagachi/kopi-santai-backend/model"
)

// MenuService is an autogenerated mock type for the MenuService type
type MenuService struct {
	mock.Mock
}

// Create provides a mock function with given fields: req
func (_m *MenuService) Create(req *dto.MenuPostReq) (*dto.MenuRes, error) {
	ret := _m.Called(req)

	var r0 *dto.MenuRes
	if rf, ok := ret.Get(0).(func(*dto.MenuPostReq) *dto.MenuRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.MenuRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.MenuPostReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: id
func (_m *MenuService) DeleteByID(id uint) (gin.H, error) {
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

// FindAll provides a mock function with given fields: q
func (_m *MenuService) FindAll(q *model.QueryParamMenu) ([]*dto.MenuRes, error) {
	ret := _m.Called(q)

	var r0 []*dto.MenuRes
	if rf, ok := ret.Get(0).(func(*model.QueryParamMenu) []*dto.MenuRes); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.MenuRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.QueryParamMenu) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllUnscoped provides a mock function with given fields: q
func (_m *MenuService) FindAllUnscoped(q *model.QueryParamMenu) ([]*dto.MenuRes, error) {
	ret := _m.Called(q)

	var r0 []*dto.MenuRes
	if rf, ok := ret.Get(0).(func(*model.QueryParamMenu) []*dto.MenuRes); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.MenuRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.QueryParamMenu) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenuDetail provides a mock function with given fields: id
func (_m *MenuService) GetMenuDetail(id uint) (*dto.MenuDetailRes, error) {
	ret := _m.Called(id)

	var r0 *dto.MenuDetailRes
	if rf, ok := ret.Get(0).(func(uint) *dto.MenuDetailRes); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.MenuDetailRes)
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

// Update provides a mock function with given fields: id, req
func (_m *MenuService) Update(id uint, req *dto.MenuUpdateReq) (*dto.MenuRes, error) {
	ret := _m.Called(id, req)

	var r0 *dto.MenuRes
	if rf, ok := ret.Get(0).(func(uint, *dto.MenuUpdateReq) *dto.MenuRes); ok {
		r0 = rf(id, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.MenuRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *dto.MenuUpdateReq) error); ok {
		r1 = rf(id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMenuService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenuService creates a new instance of MenuService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenuService(t mockConstructorTestingTNewMenuService) *MenuService {
	mock := &MenuService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
