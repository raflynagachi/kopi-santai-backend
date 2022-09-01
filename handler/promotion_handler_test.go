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

func TestPromotionHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all promotion success", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		res := []*dto.PromotionRes{{
			Name:        "HUT RI promo",
			Description: "promo nih",
			Image:       []byte("sample"),
			MinSpent:    30000,
			Coupon:      &dto.CouponRes{},
		}}
		mockService.On("FindAll").Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/promotions", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all promotion failed", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll").Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/promotions", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestPromotionHandler_FindAllUnscoped(t *testing.T) {
	t.Run("should return statusOK when find all promotion unscoped success", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		res := []*dto.PromotionRes{{
			Name:        "HUT RI promo",
			Description: "promo nih",
			Image:       []byte("sample"),
			MinSpent:    30000,
			Coupon:      &dto.CouponRes{},
		}}
		mockService.On("FindAllUnscoped").Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/internal/promotions", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all promotion unscoped failed", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAllUnscoped").Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/internal/promotions", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestPromotionHandler_CreatePromotion(t *testing.T) {
	t.Run("should return statusOK when create promotion success", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		promoReq := &dto.PromotionPostReq{
			CouponID:    1,
			Name:        "HUT RI promo",
			Description: "promo nih",
			Image:       []byte("simple"),
			MinSpent:    30000,
		}
		res := &dto.PromotionRes{
			ID:          1,
			Name:        "HUT RI promo",
			Description: "promo nih",
			Image:       []byte("sample"),
			MinSpent:    30000,
			Coupon:      &dto.CouponRes{},
		}
		mockService.On("CreatePromotion", promoReq).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		reqBody := testutils.MakeRequestBody(promoReq)
		req, _ := http.NewRequest(http.MethodPost, "/promotions", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create promotion failed", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		promoReq := &dto.PromotionPostReq{
			CouponID:    1,
			Name:        "HUT RI promo",
			Description: "promo nih",
			Image:       []byte("simple"),
			MinSpent:    30000,
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("CreatePromotion", promoReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(promoReq)
		req, _ := http.NewRequest(http.MethodPost, "/promotions", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestPromotionHandler_DeletePromotionByID(t *testing.T) {
	t.Run("should return statusOK when delete promotion by id success", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		res := gin.H{"isDeleted": true}
		mockService.On("DeletePromotionByID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodDelete, "/promotions/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when delete promotion by id failed", func(t *testing.T) {
		mockService := new(mocks.PromotionService)
		config := server.RouterConfig{PromoService: mockService}
		res := gin.H{"isDeleted": false}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("DeletePromotionByID", uint(1)).Return(res, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodDelete, "/promotions/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
