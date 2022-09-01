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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMenuHandler_FindAll(t *testing.T) {
	t.Run("should return statusOK when find all menu success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuRes := []*dto.MenuRes{{
			ID:           1,
			CategoryID:   1,
			CategoryName: "Coffee",
			Name:         "Cappuccino",
			Price:        15000,
			Image:        nil,
			Rating:       4.5,
		}}
		queryParam := &model.QueryParamMenu{
			SortBy: "id",
			Sort:   "asc",
		}
		mockService.On("FindAll", queryParam).Return(menuRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(menuRes))

		req, _ := http.NewRequest(http.MethodGet, "/menus", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all menu failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		queryParam := &model.QueryParamMenu{
			SortBy: "id",
			Sort:   "asc",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAll", queryParam).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/menus", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestMenuHandler_FindAllUnscoped(t *testing.T) {
	t.Run("should return statusOK when find all menu unscoped success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuRes := []*dto.MenuRes{{
			ID:           1,
			CategoryID:   1,
			CategoryName: "Coffee",
			Name:         "Cappuccino",
			Price:        15000,
			Image:        nil,
			Rating:       4.5,
		}}
		queryParam := &model.QueryParamMenu{
			SortBy: "id",
			Sort:   "asc",
		}
		mockService.On("FindAllUnscoped", queryParam).Return(menuRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(menuRes))

		req, _ := http.NewRequest(http.MethodGet, "/internal/menus", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find all menu unscoped failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		queryParam := &model.QueryParamMenu{
			SortBy: "id",
			Sort:   "asc",
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("FindAllUnscoped", queryParam).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/internal/menus", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestMenuHandler_GetMenuDetail(t *testing.T) {
	t.Run("should return statusOK when find menu detail success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuRes := &dto.MenuDetailRes{
			MenuRes:    &dto.MenuRes{},
			MenuOption: []*dto.MenuOptionRes{},
		}
		mockService.On("GetMenuDetail", uint(1)).Return(menuRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(menuRes))

		req, _ := http.NewRequest(http.MethodGet, "/menus/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when find menu detail failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("GetMenuDetail", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodGet, "/menus/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestMenuHandler_CreateMenu(t *testing.T) {
	t.Run("should return statusOK when create menu success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuReq := &dto.MenuPostReq{
			CategoryID: 1,
			Name:       "Cappuccino",
			Price:      15000,
			Image:      []byte("sample"),
		}
		menuRes := &dto.MenuRes{
			ID:           1,
			CategoryID:   1,
			CategoryName: "Coffee",
			Name:         "Cappuccino",
			Price:        15000,
			Image:        []byte("sample"),
			Rating:       4.5,
		}
		mockService.On("Create", menuReq).Return(menuRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(menuRes))

		reqBody := testutils.MakeRequestBody(menuReq)
		req, _ := http.NewRequest(http.MethodPost, "/menus", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when create menu failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuReq := &dto.MenuPostReq{
			CategoryID: 1,
			Name:       "Cappuccino",
			Price:      15000,
			Image:      []byte("sample"),
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("Create", menuReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(menuReq)
		req, _ := http.NewRequest(http.MethodPost, "/menus", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestMenuHandler_UpdateMenu(t *testing.T) {
	t.Run("should return statusOK when update menu success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuReq := &dto.MenuUpdateReq{
			CategoryID: 1,
			Name:       "Cappuccino",
			Price:      15000,
			Image:      []byte("sample"),
		}
		menuRes := &dto.MenuRes{
			ID:           1,
			CategoryID:   1,
			CategoryName: "Coffee",
			Name:         "Cappuccino",
			Price:        15000,
			Image:        []byte("sample"),
			Rating:       4.5,
		}
		mockService.On("Update", uint(1), menuReq).Return(menuRes, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(menuRes))

		reqBody := testutils.MakeRequestBody(menuReq)
		req, _ := http.NewRequest(http.MethodPatch, "/menus/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when update menu failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		menuReq := &dto.MenuUpdateReq{
			CategoryID: 1,
			Name:       "Cappuccino",
			Price:      15000,
			Image:      []byte("sample"),
		}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("Update", uint(1), menuReq).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		reqBody := testutils.MakeRequestBody(menuReq)
		req, _ := http.NewRequest(http.MethodPatch, "/menus/1", reqBody)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestMenuHandler_DeleteByID(t *testing.T) {
	t.Run("should return statusOK when delete menu success", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		res := gin.H{"isDeleted": true}
		mockService.On("DeleteByID", uint(1)).Return(res, nil)
		expectedBody, _ := json.Marshal(dto.StatusOKResponse(res))

		req, _ := http.NewRequest(http.MethodDelete, "/menus/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})

	t.Run("should return error when delete menu failed", func(t *testing.T) {
		mockService := new(mocks.MenuService)
		config := server.RouterConfig{MenuService: mockService}
		internalServerErr := apperror.InternalServerError(errors.New("db error").Error())
		mockService.On("DeleteByID", uint(1)).Return(nil, internalServerErr)
		expectedBody, _ := json.Marshal(internalServerErr)

		req, _ := http.NewRequest(http.MethodDelete, "/menus/1", nil)
		_, w := testutils.ServeReq(&config, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}
