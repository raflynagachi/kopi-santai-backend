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

func TestReviewHandler_Create(t *testing.T) {
	t.Run("should return statusOK when create review success", func(t *testing.T) {
		mockService := new(mocks.ReviewService)
		config := server.RouterConfig{ReviewService: mockService}
		reviewReq := &dto.ReviewPostReq{
			MenuID:      1,
			Description: "Mantap",
			Rating:      4.3,
		}
		res := &dto.ReviewRes{
			MenuID:      1,
			Description: "Mantap",
			Rating:      4.3,
		}
		mockService.On("Create", reviewReq, uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		reqBody := testutils.MakeRequestBody(reviewReq)
		req, _ := http.NewRequest(http.MethodPost, "/reviews", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create review failed", func(t *testing.T) {
		mockService := new(mocks.ReviewService)
		config := server.RouterConfig{ReviewService: mockService}
		reviewReq := &dto.ReviewPostReq{
			MenuID:      1,
			Description: "Mantap",
			Rating:      4.3,
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("Create", reviewReq, uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(reviewReq)
		req, _ := http.NewRequest(http.MethodPost, "/reviews", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestReviewHandler_FindByMenuID(t *testing.T) {
	t.Run("should return statusOK when find review by id success", func(t *testing.T) {
		mockService := new(mocks.ReviewService)
		config := server.RouterConfig{ReviewService: mockService}
		res := []*dto.ReviewRes{{
			MenuID:      1,
			Description: "Mantap",
			Rating:      4.3,
		}}
		mockService.On("FindByMenuID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/menus/1/reviews", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find review by id failed", func(t *testing.T) {
		mockService := new(mocks.ReviewService)
		config := server.RouterConfig{ReviewService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindByMenuID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/menus/1/reviews", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
