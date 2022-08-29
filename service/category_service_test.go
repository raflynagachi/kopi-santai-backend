package service_test

import (
	"errors"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper/testutils"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/mocks"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
