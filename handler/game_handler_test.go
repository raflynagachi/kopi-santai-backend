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

func TestGameHandler_FindByUserID(t *testing.T) {
	t.Run("should return statusOK when find user game score success", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		gameRes := &dto.GameLeaderboardRes{
			Name:  "John Doe",
			Score: 100,
		}
		mockService.On("FindByUserID", uint(1)).Return(gameRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(gameRes))

		req, _ := http.NewRequest(http.MethodGet, "/games/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when userID and paramID not same", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindByUserID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))

		req, _ := http.NewRequest(http.MethodGet, "/games/2", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find user game score failed", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindByUserID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/games/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestGameHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all game score success", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		gameRes := []*dto.GameLeaderboardRes{{
			Name:  "John Doe",
			Score: 100,
		}, {
			Name:  "Alice Doe",
			Score: 80,
		}}
		mockService.On("FindAll").Return(gameRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(gameRes))

		req, _ := http.NewRequest(http.MethodGet, "/games", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all game score failed", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll").Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/games", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestGameHandler_AddCouponPrizeToUser(t *testing.T) {
	t.Run("should return statusOK when calculated score with prize success", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		userCouponRes := &dto.UserCouponRes{
			UserRes:   &dto.UserRes{},
			CouponRes: &dto.CouponRes{},
		}
		gameReq := &dto.GameResultPostReq{Score: 10}
		mockService.On("AddCouponPrizeToUser", gameReq, uint(1)).Return(userCouponRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(userCouponRes))

		reqBody := testutils.MakeRequestBody(gameReq)
		req, _ := http.NewRequest(http.MethodPost, "/game-prize", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when calculated score with prize failed", func(t *testing.T) {
		mockService := new(mocks.GameService)
		config := server.RouterConfig{GameService: mockService}
		gameReq := &dto.GameResultPostReq{Score: 10}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("AddCouponPrizeToUser", gameReq, uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(gameReq)
		req, _ := http.NewRequest(http.MethodPost, "/game-prize", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
