package handler_test

import (
	"encoding/json"
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestDeliveryHandler_UpdateStatus(t *testing.T) {
	t.Run("should return statusOK when update delivery success", func(t *testing.T) {
		mockService := new(mocks.DeliveryService)
		config := server.RouterConfig{DeliveryService: mockService}
		deliveryReq := &dto.DeliveryUpdateStatusReq{
			Status: "DELIVERED",
		}
		deliveryRes := &dto.DeliveryRes{
			DeliveryDate: time.Now(),
			Status:       "DELIVERED",
		}
		mockService.On("UpdateStatus", uint(1), deliveryReq).Return(deliveryRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(deliveryRes))

		reqBody := testutils.MakeRequestBody(deliveryReq)
		req, _ := http.NewRequest(http.MethodPatch, "/deliveries/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when update delivery failed", func(t *testing.T) {
		mockService := new(mocks.DeliveryService)
		config := server.RouterConfig{DeliveryService: mockService}
		deliveryReq := &dto.DeliveryUpdateStatusReq{
			Status: "DELIVERED",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("UpdateStatus", uint(1), deliveryReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(deliveryReq)
		req, _ := http.NewRequest(http.MethodPatch, "/deliveries/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
