package handler

import (
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
	FindAllUnscoped(c *gin.Context)
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

func (h *menuHandler) FindAllUnscoped(c *gin.Context) {
	queryParam := &model.QueryParamMenu{
		Search:   helper.GetQuery(c, "search", ""),
		SortBy:   helper.GetQuery(c, "sortBy", sortMenuByDefaultValue),
		Sort:     helper.GetQuery(c, "sort", sortMenuDefaultValue),
		Category: helper.GetQuery(c, "category", ""),
	}

	menusRes, err := h.menuService.FindAllUnscoped(queryParam)
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
	payload, _ := c.Get("payload")
	var req *dto.MenuPostReq
	req = payload.(*dto.MenuPostReq)

	menuRes, err := h.menuService.Create(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(menuRes))
}

func (h *menuHandler) UpdateMenu(c *gin.Context) {
	idParam, _ := c.Get("id")

	payload, _ := c.Get("payload")
	var req *dto.MenuUpdateReq
	req = payload.(*dto.MenuUpdateReq)

	menuRes, err := h.menuService.Update(idParam.(uint), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(menuRes))
}

func (h *menuHandler) DeleteByID(c *gin.Context) {
	idParam, _ := c.Get("id")

	res, err := h.menuService.DeleteByID(idParam.(uint))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.StatusOKResponse(res))
}
