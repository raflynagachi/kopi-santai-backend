package service_test

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
