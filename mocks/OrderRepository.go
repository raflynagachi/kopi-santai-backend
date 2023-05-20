// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/raflynagachi/kopi-santai-backend/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	time "time"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// CountRecords provides a mock function with given fields: tx, t
func (_m *OrderRepository) CountRecords(tx *gorm.DB, t *time.Time) (int, error) {
	ret := _m.Called(tx, t)

	var r0 int
	if rf, ok := ret.Get(0).(func(*gorm.DB, *time.Time) int); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *time.Time) error); ok {
		r1 = rf(tx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: tx, o
func (_m *OrderRepository) CreateOrder(tx *gorm.DB, o *model.Order) (*model.Order, error) {
	ret := _m.Called(tx, o)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(*gorm.DB, *model.Order) *model.Order); ok {
		r0 = rf(tx, o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *model.Order) error); ok {
		r1 = rf(tx, o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: tx, t, limit, page
func (_m *OrderRepository) FindAll(tx *gorm.DB, t *time.Time, limit int, page int) ([]*model.Order, error) {
	ret := _m.Called(tx, t, limit, page)

	var r0 []*model.Order
	if rf, ok := ret.Get(0).(func(*gorm.DB, *time.Time, int, int) []*model.Order); ok {
		r0 = rf(tx, t, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *time.Time, int, int) error); ok {
		r1 = rf(tx, t, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOrderByIDAndUserID provides a mock function with given fields: tx, id, userID
func (_m *OrderRepository) FindOrderByIDAndUserID(tx *gorm.DB, id uint, userID uint) (*model.Order, error) {
	ret := _m.Called(tx, id, userID)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint, uint) *model.Order); ok {
		r0 = rf(tx, id, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
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

// FindOrderByUserID provides a mock function with given fields: tx, userID
func (_m *OrderRepository) FindOrderByUserID(tx *gorm.DB, userID uint) ([]*model.Order, error) {
	ret := _m.Called(tx, userID)

	var r0 []*model.Order
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) []*model.Order); ok {
		r0 = rf(tx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
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

// SumOfTotalPrice provides a mock function with given fields: tx, t
func (_m *OrderRepository) SumOfTotalPrice(tx *gorm.DB, t *time.Time) (float64, error) {
	ret := _m.Called(tx, t)

	var r0 float64
	if rf, ok := ret.Get(0).(func(*gorm.DB, *time.Time) float64); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *time.Time) error); ok {
		r1 = rf(tx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: tx, id, ord
func (_m *OrderRepository) Update(tx *gorm.DB, id uint, ord *model.Order) (*model.Order, error) {
	ret := _m.Called(tx, id, ord)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint, *model.Order) *model.Order); ok {
		r0 = rf(tx, id, ord)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint, *model.Order) error); ok {
		r1 = rf(tx, id, ord)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderRepository(t mockConstructorTestingTNewOrderRepository) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
