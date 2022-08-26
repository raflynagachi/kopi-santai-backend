package handler_test

import (
	"encoding/json"
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCouponHandler_Create(t *testing.T) {
	t.Run("should return statusOK when create success", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		payload := dto.CouponPostReq{
			Name:   "Special Discount HUT Kopi Santai",
			Amount: 30,
		}
		couponRes := dto.CouponRes{
			ID:          1,
			Name:        payload.Name,
			Amount:      payload.Amount,
			IsAvailable: true,
		}
		mockService.On("Create", &payload).Return(&couponRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(couponRes))

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/coupons", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create failed", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		payload := dto.CouponPostReq{
			Name:   "Special Discount HUT Kopi Santai",
			Amount: 30,
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("Create", &payload).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/coupons", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestCouponHandler_FindCouponByUserID(t *testing.T) {
	t.Run("should return statusOK when find coupon success", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		couponRes := []*dto.CouponRes{{
			ID:          1,
			Name:        "Special Discount HUT Kopi Santai",
			Amount:      30,
			IsAvailable: true,
		}}
		userID := uint(1)
		mockService.On("FindCouponByUserID", userID).Return(couponRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(couponRes))

		req, _ := http.NewRequest(http.MethodGet, "/coupons", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find coupon failed", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		userID := uint(1)
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindCouponByUserID", userID).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/coupons", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestCouponHandler_DeleteByID(t *testing.T) {
	t.Run("should return statusOK when delete coupon success", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		res := gin.H{"isDeleted": true}
		paramID := uint(1)
		mockService.On("DeleteByID", paramID).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodDelete, "/coupons/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when delete coupon failed", func(t *testing.T) {
		mockService := new(mocks.CouponService)
		config := server.RouterConfig{CouponService: mockService}
		paramID := uint(1)
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("DeleteByID", paramID).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodDelete, "/coupons/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
