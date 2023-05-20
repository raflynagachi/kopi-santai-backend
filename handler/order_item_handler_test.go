package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/server"
	"github.com/stretchr/testify/assert"
)

func TestOrderItemHandler_CreateOrderItem(t *testing.T) {
	t.Run("should return statusOK when create order item success", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		orderItemReq := &dto.OrderItemPostReq{
			MenuID:      1,
			Quantity:    2,
			Description: "toppings:boba",
		}
		res := &dto.OrderItemRes{
			Menu:        &dto.MenuRes{},
			Quantity:    2,
			Description: "toppings:boba",
		}
		mockService.On("CreateOrderItem", orderItemReq, uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		reqBody := testutils.MakeRequestBody(orderItemReq)
		req, _ := http.NewRequest(http.MethodPost, "/order-items", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create order item failed", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		orderItemReq := &dto.OrderItemPostReq{
			MenuID:      1,
			Quantity:    2,
			Description: "toppings:boba",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("CreateOrderItem", orderItemReq, uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(orderItemReq)
		req, _ := http.NewRequest(http.MethodPost, "/order-items", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderItemHandler_FindOrderItemByUserID(t *testing.T) {
	t.Run("should return statusOK when find order item by user id success", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		res := []*dto.OrderItemRes{{
			Menu:        &dto.MenuRes{},
			Quantity:    2,
			Description: "toppings:boba",
		}}
		mockService.On("FindOrderItemByUserID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/order-items", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find order item by user id failed", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindOrderItemByUserID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/order-items", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderItemHandler_UpdateOrderItemByID(t *testing.T) {
	t.Run("should return statusOK when update order item by id success", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		orderItemReq := &dto.OrderItemPatchReq{
			Quantity:    3,
			Description: "toppings:boba",
		}
		res := &dto.OrderItemRes{
			Menu:        &dto.MenuRes{},
			Quantity:    3,
			Description: "toppings:boba",
		}
		mockService.On("UpdateOrderItemByID", uint(1), uint(1), orderItemReq).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		reqBody := testutils.MakeRequestBody(orderItemReq)
		req, _ := http.NewRequest(http.MethodPatch, "/order-items/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when update order item by id failed", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		orderItemReq := &dto.OrderItemPatchReq{
			Quantity:    3,
			Description: "toppings:boba",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("UpdateOrderItemByID", uint(1), uint(1), orderItemReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(orderItemReq)
		req, _ := http.NewRequest(http.MethodPatch, "/order-items/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderItemHandler_DeleteOrderItemByID(t *testing.T) {
	t.Run("should return statusOK when delete order item by id success", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		res := gin.H{"isDeleted": true}
		mockService.On("DeleteOrderItemByID", uint(1), uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodDelete, "/order-items/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when delete order item by id failed", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("DeleteOrderItemByID", uint(1), uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodDelete, "/order-items/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestOrderItemHandler_DeleteOrderItemByUserID(t *testing.T) {
	t.Run("should return statusOK when delete order item by userID success", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		res := gin.H{"isDeleted": true}
		mockService.On("DeleteOrderItemByUserID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodDelete, "/order-items", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when delete order item by userID failed", func(t *testing.T) {
		mockService := new(mocks.OrderItemService)
		config := server.RouterConfig{OrderItemService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("DeleteOrderItemByUserID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodDelete, "/order-items", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
