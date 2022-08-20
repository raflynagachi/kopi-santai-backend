package handler

import (
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/helper"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/model"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const sortDefaultValue = model.Asc
const sortByDefaultValue = model.MenuID

type MenuHandler interface {
	FindAll(c *gin.Context)
	GetMenuDetail(c *gin.Context)
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
		SortBy:   helper.GetQuery(c, "sortBy", sortByDefaultValue),
		Sort:     helper.GetQuery(c, "sort", sortDefaultValue),
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
