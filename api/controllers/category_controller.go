package controllers

import (
	"clean-architecture-api/api/responses"
	"clean-architecture-api/lib"
	"clean-architecture-api/services"
	"clean-architecture-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *services.CategoryService
	logger  lib.Logger
}

func NewCategoryController(categoryService *services.CategoryService, logger lib.Logger) *CategoryController {
	return &CategoryController{
		service: categoryService,
		logger:  logger,
	}
}

func (ctrl *CategoryController) GetCategory(c *gin.Context) {
	categories, err := ctrl.service.SetPaginationScope(utils.Paginate(c)).GetAll()

	if err != nil {
		ctrl.logger.Error(err)
	}
	responses.JSONWithPagination(c, http.StatusOK, categories)
}
