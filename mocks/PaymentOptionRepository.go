// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/raflynagachi/kopi-santai-backend/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// PaymentOptionRepository is an autogenerated mock type for the PaymentOptionRepository type
type PaymentOptionRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: tx
func (_m *PaymentOptionRepository) FindAll(tx *gorm.DB) ([]*model.PaymentOption, error) {
	ret := _m.Called(tx)

	var r0 []*model.PaymentOption
	if rf, ok := ret.Get(0).(func(*gorm.DB) []*model.PaymentOption); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PaymentOption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: tx, id
func (_m *PaymentOptionRepository) FindByID(tx *gorm.DB, id uint) (*model.PaymentOption, error) {
	ret := _m.Called(tx, id)

	var r0 *model.PaymentOption
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) *model.PaymentOption); ok {
		r0 = rf(tx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PaymentOption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPaymentOptionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPaymentOptionRepository creates a new instance of PaymentOptionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPaymentOptionRepository(t mockConstructorTestingTNewPaymentOptionRepository) *PaymentOptionRepository {
	mock := &PaymentOptionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
