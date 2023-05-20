package service_test

import (
	"errors"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPaymentOptionService_FindAll(t *testing.T) {
	t.Run("should return response when find all payment options success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PaymentOptionRepository)
		cfg := &service.PaymentOptConfig{
			DB:             gormDB,
			PaymentOptRepo: mockRepository,
		}
		s := service.NewPaymentOpt(cfg)
		expectedRes := []*dto.PaymentOptionRes{{
			ID:   1,
			Name: "ShopeePay",
		}}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.PaymentOption{{
			ID:   1,
			Name: "ShopeePay",
		}}, nil)

		menuRes, err := s.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find all payment options failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.PaymentOptionRepository)
		cfg := &service.PaymentOptConfig{
			DB:             gormDB,
			PaymentOptRepo: mockRepository,
		}
		s := service.NewPaymentOpt(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
