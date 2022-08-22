package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/apperror"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const sortMenuDefaultValue = model.Asc
const sortMenuByDefaultValue = model.ID

type MenuHandler interface {
	FindAll(c *gin.Context)
	GetMenuDetail(c *gin.Context)
	CreateMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type menuHandler struct {
	menuService service.MenuService
}

type MenuConfig struct {
	MenuService service.MenuService
}

func NewMenu(c *MenuConfig) MenuHandler {
	return &menuHandler{
		menuService: c.MenuService,
	}
}

func (h *menuHandler) FindAll(c *gin.Context) {
	queryParam := &model.QueryParamMenu{
		Search:   helper.GetQuery(c, "search", ""),
		SortBy:   helper.GetQuery(c, "sortBy", sortMenuByDefaultValue),
		Sort:     helper.GetQuery(c, "sort", sortMenuDefaultValue),
		Category: helper.GetQuery(c, "category", ""),
	}

	menusRes, err := h.menuService.FindAll(queryParam)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(menusRes))
}

func (h *menuHandler) GetMenuDetail(c *gin.Context) {
	idParam, _ := c.Get("id")
	menuRes, err := h.menuService.GetMenuDetail(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(menuRes))
}

func (h *menuHandler) CreateMenu(c *gin.Context) {
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	if role != model.AdminRole {
		_ = c.Error(apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error()))
		return
	}

	payload, _ := c.Get("payload")
	var req *dto.MenuPostReq
	req = payload.(*dto.MenuPostReq)

	orderItemRes, err := h.menuService.Create(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *menuHandler) UpdateMenu(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	if role != model.AdminRole {
		_ = c.Error(apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error()))
		return
	}

	payload, _ := c.Get("payload")
	var req *dto.MenuUpdateReq
	req = payload.(*dto.MenuUpdateReq)

	orderItemRes, err := h.menuService.Update(idParam.(uint), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}

func (h *menuHandler) DeleteByID(c *gin.Context) {
	idParam, _ := c.Get("id")
	userPayload, ok := c.Get("user")
	if !ok {
		_ = c.Error(apperror.UnauthorizedError(new(apperror.UserUnauthorizedError).Error()))
		return
	}
	role := userPayload.(*dto.UserJWT).Role
	if role != model.AdminRole {
		_ = c.Error(apperror.ForbiddenError(new(apperror.ForbiddenAccessError).Error()))
		return
	}

	orderItemRes, err := h.menuService.DeleteByID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(orderItemRes))
}
