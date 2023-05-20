// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/raflynagachi/kopi-santai-backend/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// CouponRepository is an autogenerated mock type for the CouponRepository type
type CouponRepository struct {
	mock.Mock
}

// AddCouponToUser provides a mock function with given fields: tx, uc
func (_m *CouponRepository) AddCouponToUser(tx *gorm.DB, uc *model.UserCoupon) (*model.UserCoupon, error) {
	ret := _m.Called(tx, uc)

	var r0 *model.UserCoupon
	if rf, ok := ret.Get(0).(func(*gorm.DB, *model.UserCoupon) *model.UserCoupon); ok {
		r0 = rf(tx, uc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserCoupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *model.UserCoupon) error); ok {
		r1 = rf(tx, uc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: tx, c
func (_m *CouponRepository) Create(tx *gorm.DB, c *model.Coupon) (*model.Coupon, error) {
	ret := _m.Called(tx, c)

	var r0 *model.Coupon
	if rf, ok := ret.Get(0).(func(*gorm.DB, *model.Coupon) *model.Coupon); ok {
		r0 = rf(tx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Coupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *model.Coupon) error); ok {
		r1 = rf(tx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: tx, id
func (_m *CouponRepository) DeleteByID(tx *gorm.DB, id uint) (bool, error) {
	ret := _m.Called(tx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) bool); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserCoupon provides a mock function with given fields: tx, id
func (_m *CouponRepository) DeleteUserCoupon(tx *gorm.DB, id uint) (bool, error) {
	ret := _m.Called(tx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) bool); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserCouponByCouponID provides a mock function with given fields: tx, couponID
func (_m *CouponRepository) DeleteUserCouponByCouponID(tx *gorm.DB, couponID uint) (bool, error) {
	ret := _m.Called(tx, couponID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) bool); ok {
		r0 = rf(tx, couponID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, couponID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: tx
func (_m *CouponRepository) FindAll(tx *gorm.DB) ([]*model.Coupon, error) {
	ret := _m.Called(tx)

	var r0 []*model.Coupon
	if rf, ok := ret.Get(0).(func(*gorm.DB) []*model.Coupon); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Coupon)
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
func (_m *CouponRepository) FindByID(tx *gorm.DB, id uint) (*model.Coupon, error) {
	ret := _m.Called(tx, id)

	var r0 *model.Coupon
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) *model.Coupon); ok {
		r0 = rf(tx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Coupon)
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

// FindCouponByUserID provides a mock function with given fields: tx, userID
func (_m *CouponRepository) FindCouponByUserID(tx *gorm.DB, userID uint) ([]*model.UserCoupon, error) {
	ret := _m.Called(tx, userID)

	var r0 []*model.UserCoupon
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) []*model.UserCoupon); ok {
		r0 = rf(tx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserCoupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserCouponByCouponID provides a mock function with given fields: tx, id, userID
func (_m *CouponRepository) FindUserCouponByCouponID(tx *gorm.DB, id uint, userID uint) (*model.UserCoupon, error) {
	ret := _m.Called(tx, id, userID)

	var r0 *model.UserCoupon
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint, uint) *model.UserCoupon); ok {
		r0 = rf(tx, id, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserCoupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint, uint) error); ok {
		r1 = rf(tx, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCouponRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCouponRepository creates a new instance of CouponRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCouponRepository(t mockConstructorTestingTNewCouponRepository) *CouponRepository {
	mock := &CouponRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
