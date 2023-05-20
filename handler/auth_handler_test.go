package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/server"
	"github.com/stretchr/testify/assert"
)

var tokenResp = dto.TokenRes{Token: "jwt_generated_token"}

func TestAuthHandler_Register(t *testing.T) {
	t.Run("should return statusOK, user, and idToken when register success", func(t *testing.T) {
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		payload := dto.RegisterPostReq{
			FullName: "John Doe",
			Phone:    "+6282212345678",
			Address:  "Jl. Kemerdekaan, Nusa Bangsa",
			Email:    "john@mail.com",
			Password: "password",
		}
		mockService.On("Register", &payload).Return(&tokenResp, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(tokenResp))

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/register", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return http error when register failed", func(t *testing.T) {
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		payload := dto.RegisterPostReq{
			FullName: "John Doe",
			Phone:    "+6282212345678",
			Address:  "Jl. Kemerdekaan, Nusa Bangsa",
			Email:    "john@mail.com",
			Password: "password",
		}
		internalServerErr := apperror.InternalServerError(new(apperror.EmailAlreadyExistError).Error())
		mockService.On("Register", &payload).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/register", requestBody)
		_, w := testutils.ServeReq(&config, req)

		fmt.Println(w.Body.String())
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("should return statusOK and idToken when login success", func(t *testing.T) {
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		payload := dto.LoginPostReq{
			Email:    "john@mail.com",
			Password: "password",
		}
		mockService.On("Login", &payload).Return(&tokenResp, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(tokenResp))

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/login", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return http error when login failed", func(t *testing.T) {
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		payload := dto.LoginPostReq{
			Email:    "john@mail.com",
			Password: "password",
		}
		unauthorizedErr := apperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Status:     "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
		mockService.On("Login", &payload).Return(nil, unauthorizedErr)
		expectedBody, _ := json.Marshal(unauthorizedErr)

		requestBody := testutils.MakeRequestBody(payload)
		req, _ := http.NewRequest(http.MethodPost, "/login", requestBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
