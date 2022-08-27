package service_test

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var menu = &model.Menu{
	Category: &model.Category{
		MenuOptionsCategories: &model.MenuOptionsCategories{
			MenuOption: &model.MenuOption{},
		},
	},
	Reviews: []*model.Review{{}},
}

func TestMenuService_FindAll(t *testing.T) {
	t.Run("should return response when find all menu success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		queryParam := &model.QueryParamMenu{
			Search:   "co",
			SortBy:   "price",
			Sort:     "asc",
			Category: "coffee",
		}
		expectedRes := []*dto.MenuRes{{}}
		output := []*model.Menu{menu}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), queryParam).Return(output, nil)

		menuRes, err := s.FindAll(queryParam)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find all menu failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		queryParam := &model.QueryParamMenu{}
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType), queryParam).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll(queryParam)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestMenuService_GetMenuDetail(t *testing.T) {
	t.Run("should return response when get menu detail success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		expectedRes := &dto.MenuDetailRes{
			MenuRes:    &dto.MenuRes{},
			MenuOption: []*dto.MenuOptionRes{{}},
		}
		output := []*model.MenuOptionsCategories{{
			MenuOption: &model.MenuOption{},
		}}
		mockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(menu, nil)
		mockRepository.On("FindMenuOptions", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(output, nil)

		menuRes, err := s.GetMenuDetail(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when menu detail not found", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(nil, dbErr)
		expectedErr := apperror.NotFoundError(dbErr.Error())

		_, err := s.GetMenuDetail(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})

	t.Run("should return error when failed to get menu options", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(menu, nil)
		mockRepository.On("FindMenuOptions", mock.AnythingOfType(testutils.GormDBPointerType), uint(0)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.GetMenuDetail(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestMenuService_Create(t *testing.T) {
	t.Run("should return response when create menu success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		req := &dto.MenuPostReq{}
		expectedRes := &dto.MenuRes{}
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Menu")).Return(menu, nil)

		menuRes, err := s.Create(req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when create menu failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		req := &dto.MenuPostReq{}
		dbErr := errors.New("db error")
		mockRepository.On("Create", mock.AnythingOfType(testutils.GormDBPointerType), mock.AnythingOfType("*model.Menu")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.Create(req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestMenuService_Update(t *testing.T) {
	t.Run("should return response when update menu success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		req := &dto.MenuUpdateReq{}
		expectedRes := &dto.MenuRes{}
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.Menu")).Return(menu, nil)

		menuRes, err := s.Update(uint(1), req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when update menu failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		req := &dto.MenuUpdateReq{}
		dbErr := errors.New("db error")
		mockRepository.On("Update", mock.AnythingOfType(testutils.GormDBPointerType), uint(1), mock.AnythingOfType("*model.Menu")).Return(nil, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.Update(uint(1), req)

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}

func TestMenuService_DeleteByID(t *testing.T) {
	t.Run("should return response when delete menu by id success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		expectedRes := gin.H{"isDeleted": true}
		mockRepository.On("DeleteByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(true, nil)

		menuRes, err := s.DeleteByID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when delete menu by id failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.MenuRepository)
		cfg := &service.MenuConfig{
			DB:             gormDB,
			MenuRepository: mockRepository,
		}
		s := service.NewMenu(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("DeleteByID", mock.AnythingOfType(testutils.GormDBPointerType), uint(1)).Return(false, dbErr)
		expectedErr := apperror.BadRequestError(dbErr.Error())

		_, err := s.DeleteByID(uint(1))

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
