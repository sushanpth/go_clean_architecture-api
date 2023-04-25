package controllers

import (
	"clean-architecture-api/api/responses"
	"clean-architecture-api/lib"
	"clean-architecture-api/models"
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

func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBind(&category); err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed get data from request!")
		return
	}

	if err := ctrl.service.Create(category); err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed create new category!")
		return
	}

	c.JSON(200, gin.H{"data": "category created"})

}
