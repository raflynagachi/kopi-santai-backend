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

func TestCategoryHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all category success", func(t *testing.T) {
		mockService := new(mocks.CategoryService)
		config := server.RouterConfig{CategoryService: mockService}
		res := []*dto.CategoryRes{{
			ID:   1,
			Name: "Coffee",
		}}
		mockService.On("FindAll").Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all category failed", func(t *testing.T) {
		mockService := new(mocks.CategoryService)
		config := server.RouterConfig{CategoryService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll").Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
