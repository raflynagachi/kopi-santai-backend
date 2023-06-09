// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/raflynagachi/kopi-santai-backend/dto"
	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Login provides a mock function with given fields: req
func (_m *AuthService) Login(req *dto.LoginPostReq) (*dto.TokenRes, error) {
	ret := _m.Called(req)

	var r0 *dto.TokenRes
	if rf, ok := ret.Get(0).(func(*dto.LoginPostReq) *dto.TokenRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.TokenRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.LoginPostReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: req
func (_m *AuthService) Register(req *dto.RegisterPostReq) (*dto.TokenRes, error) {
	ret := _m.Called(req)

	var r0 *dto.TokenRes
	if rf, ok := ret.Get(0).(func(*dto.RegisterPostReq) *dto.TokenRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.TokenRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.RegisterPostReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthService(t mockConstructorTestingTNewAuthService) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
