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

var delivery = model.Delivery{
	ID:     1,
	Status: model.StatusDefault,
}

func TestDeliveryService_UpdateStatus(t *testing.T) {
	t.Run("should return response when update status delivery success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.DeliveryRepository)
		deliveryConfig := &service.DeliveryConfig{DB: gormDB, DeliveryRepo: mockRepository}
		s := service.NewDelivery(deliveryConfig)
		req := &dto.DeliveryUpdateStatusReq{
			Status: model.StatusDefault,
		}
		expectedRes := &dto.DeliveryRes{
			Status: model.StatusDefault,
		}
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.Delivery")).Return(&delivery, nil)

		actualRes, err := s.UpdateStatus(uint(1), req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("should return error when update status delivery failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.DeliveryRepository)
		deliveryConfig := &service.DeliveryConfig{DB: gormDB, DeliveryRepo: mockRepository}
		s := service.NewDelivery(deliveryConfig)
		req := &dto.DeliveryUpdateStatusReq{
			Status: "ON PROCESS",
		}
		dbErr := errors.New("db error")
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.Delivery")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.UpdateStatus(uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
