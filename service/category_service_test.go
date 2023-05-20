package service_test

import (
	"errors"
	"testing"

	"github.com/raflynagachi/kopi-santai-backend/apperror"
	"github.com/raflynagachi/kopi-santai-backend/dto"
	"github.com/raflynagachi/kopi-santai-backend/helper/testutils"
	"github.com/raflynagachi/kopi-santai-backend/mocks"
	"github.com/raflynagachi/kopi-santai-backend/model"
	"github.com/raflynagachi/kopi-santai-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCategoryService_FindAll(t *testing.T) {
	t.Run("should return response when find all categories success", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CategoryRepository)
		cfg := &service.CategoryConfig{
			DB:           gormDB,
			CategoryRepo: mockRepository,
		}
		s := service.NewCategory(cfg)
		expectedRes := []*dto.CategoryRes{{}}
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return([]*model.Category{{}}, nil)

		menuRes, err := s.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, menuRes)
	})

	t.Run("should return error when find all categories failed", func(t *testing.T) {
		gormDB := testutils.MockDB()
		mockRepository := new(mocks.CategoryRepository)
		cfg := &service.CategoryConfig{
			DB:           gormDB,
			CategoryRepo: mockRepository,
		}
		s := service.NewCategory(cfg)
		dbErr := errors.New("db error")
		mockRepository.On("FindAll", mock.AnythingOfType(testutils.GormDBPointerType)).Return(nil, dbErr)
		expectedErr := apperror.InternalServerError(dbErr.Error())

		_, err := s.FindAll()

		assert.Equal(t, expectedErr, err)
		assert.ErrorContains(t, err, expectedErr.Error())
	})
}
