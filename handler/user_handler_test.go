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
)

func TestUserHandler_GetProfileDetail(t *testing.T) {
	t.Run("should return statusOK when find user detail success", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		userRes := &dto.ProfileRes{
			User:    &dto.UserRes{},
			Coupons: []*dto.CouponRes{},
		}
		mockService.On("GetProfileDetail", uint(1)).Return(userRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(userRes))

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when userID and paramID not same", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("GetProfileDetail", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))

		req, _ := http.NewRequest(http.MethodGet, "/users/2", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find user detail failed", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("GetProfileDetail", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	t.Run("should return statusOK when update user profile success", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		userReq := &dto.UserUpdateReq{
			Email: "john.doeee@mail.com",
		}
		userRes := &dto.UserRes{
			Email: "john.doee@mail.com",
		}
		mockService.On("UpdateProfile", uint(1), userReq).Return(userRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(userRes))

		reqBody := testutils.MakeRequestBody(userReq)
		req, _ := http.NewRequest(http.MethodPatch, "/users/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when userID and paramID not same", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		userReq := &dto.UserUpdateReq{
			Email: "john.doeee@mail.com",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("UpdateProfile", uint(1), userReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))

		reqBody := testutils.MakeRequestBody(userReq)
		req, _ := http.NewRequest(http.MethodPatch, "/users/2", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when update user profile failed", func(t *testing.T) {
		mockService := new(mocks.UserService)
		config := server.RouterConfig{UserService: mockService}
		userReq := &dto.UserUpdateReq{
			Email: "john.doeee@mail.com",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("UpdateProfile", uint(1), userReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(userReq)
		req, _ := http.NewRequest(http.MethodPatch, "/users/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
