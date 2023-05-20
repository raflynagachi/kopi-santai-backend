package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/server"
	"github.com/stretchr/testify/assert"
)

func TestPaymentOptionHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all success", func(t *testing.T) {
		mockService := new(mocks.PaymentOptionService)
		config := server.RouterConfig{PaymentOptService: mockService}
		res := []*dto.PaymentOptionRes{{
			ID:   1,
			Name: "ShopeePay",
		}}
		mockService.On("FindAll").Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/payment-options", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all failed", func(t *testing.T) {
		mockService := new(mocks.PaymentOptionService)
		config := server.RouterConfig{PaymentOptService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll").Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/payment-options", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
