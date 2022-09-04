package handler_test

import (
	"encoding/json"
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestOrderHandler_CreateOrder(t *testing.T) {
	t.Run("should return statusOK when create order success", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		orderReq := &dto.OrderPostReq{
			PaymentOptID: 1,
			CouponID:     1,
			OrderedDate:  time.Time{}.AddDate(1, 0, 0),
		}
		res := &dto.OrderRes{
			UserID:        1,
			CouponID:      1,
			OrderedDate:   time.Time{}.AddDate(1, 0, 0),
			TotalPrice:    100,
			Coupon:        &dto.CouponRes{},
			Delivery:      &dto.DeliveryRes{},
			PaymentOption: &dto.PaymentOptionRes{},
			OrderItems:    []*dto.OrderItemRes{},
		}
		mockService.On("CreateOrder", orderReq, uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		reqBody := testutils.MakeRequestBody(orderReq)
		req, _ := http.NewRequest(http.MethodPost, "/orders", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create order failed", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		orderReq := &dto.OrderPostReq{
			PaymentOptID: 1,
			CouponID:     1,
			OrderedDate:  time.Time{}.AddDate(1, 0, 0),
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("CreateOrder", orderReq, uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(orderReq)
		req, _ := http.NewRequest(http.MethodPost, "/orders", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderHandler_FindOrderByIDAndUserID(t *testing.T) {
	t.Run("should return statusOK when find order by id and userID success", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		res := &dto.OrderRes{
			UserID:        1,
			CouponID:      1,
			OrderedDate:   time.Time{}.AddDate(1, 0, 0),
			TotalPrice:    100,
			Coupon:        &dto.CouponRes{},
			Delivery:      &dto.DeliveryRes{},
			PaymentOption: &dto.PaymentOptionRes{},
			OrderItems:    []*dto.OrderItemRes{},
		}
		mockService.On("FindOrderByIDAndUserID", uint(1), uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find order by id and userID failed", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindOrderByIDAndUserID", uint(1), uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all order success", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		orderRes := []*dto.OrderRes{{
			UserID:        1,
			CouponID:      1,
			OrderedDate:   time.Time{}.AddDate(1, 0, 0),
			TotalPrice:    100,
			Coupon:        &dto.CouponRes{},
			Delivery:      &dto.DeliveryRes{},
			PaymentOption: &dto.PaymentOptionRes{},
			OrderItems:    []*dto.OrderItemRes{},
		}}
		res := &dto.OrderPaginationRes{
			CurrentPage: 1,
			TotalPage:   1,
			TotalData:   1,
			Limit:       10,
			OrderRes:    orderRes,
		}
		queryParam := &model.QueryParamOrder{
			Date:  "lastWeek",
			Limit: 10,
			Page:  1,
		}
		mockService.On("FindAll", queryParam).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/internal/orders?date=lastWeek&limit=10&page=1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all order failed", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		queryParam := &model.QueryParamOrder{Date: "lastWeek", Limit: 10, Page: 1}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll", queryParam).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/internal/orders?date=lastWeek", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderHandler_FindByUserID(t *testing.T) {
	t.Run("should return statusOK when find order by userID success", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		res := []*dto.OrderRes{{
			UserID:        1,
			CouponID:      1,
			OrderedDate:   time.Time{}.AddDate(1, 0, 0),
			TotalPrice:    100,
			Coupon:        &dto.CouponRes{},
			Delivery:      &dto.DeliveryRes{},
			PaymentOption: &dto.PaymentOptionRes{},
			OrderItems:    []*dto.OrderItemRes{},
		}}
		mockService.On("FindOrderByUserID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find order by userID failed", func(t *testing.T) {
		mockService := new(mocks.OrderService)
		config := server.RouterConfig{OrderService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindOrderByUserID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
