// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/raflynagachi/kopi-santai-backend/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// PromotionRepository is an autogenerated mock type for the PromotionRepository type
type PromotionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: tx, p
func (_m *PromotionRepository) Create(tx *gorm.DB, p *model.Promotion) (*model.Promotion, error) {
	ret := _m.Called(tx, p)

	var r0 *model.Promotion
	if rf, ok := ret.Get(0).(func(*gorm.DB, *model.Promotion) *model.Promotion); ok {
		r0 = rf(tx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Promotion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *model.Promotion) error); ok {
		r1 = rf(tx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: tx, id
func (_m *PromotionRepository) Delete(tx *gorm.DB, id uint) (bool, error) {
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

// FindAll provides a mock function with given fields: tx
func (_m *PromotionRepository) FindAll(tx *gorm.DB) ([]*model.Promotion, error) {
	ret := _m.Called(tx)

	var r0 []*model.Promotion
	if rf, ok := ret.Get(0).(func(*gorm.DB) []*model.Promotion); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Promotion)
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

// FindAllUnscoped provides a mock function with given fields: tx
func (_m *PromotionRepository) FindAllUnscoped(tx *gorm.DB) ([]*model.Promotion, error) {
	ret := _m.Called(tx)

	var r0 []*model.Promotion
	if rf, ok := ret.Get(0).(func(*gorm.DB) []*model.Promotion); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Promotion)
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

// FindByMinSpent provides a mock function with given fields: tx, spent
func (_m *PromotionRepository) FindByMinSpent(tx *gorm.DB, spent uint) ([]*model.Promotion, error) {
	ret := _m.Called(tx, spent)

	var r0 []*model.Promotion
	if rf, ok := ret.Get(0).(func(*gorm.DB, uint) []*model.Promotion); ok {
		r0 = rf(tx, spent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Promotion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, uint) error); ok {
		r1 = rf(tx, spent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPromotionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromotionRepository creates a new instance of PromotionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromotionRepository(t mockConstructorTestingTNewPromotionRepository) *PromotionRepository {
	mock := &PromotionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
