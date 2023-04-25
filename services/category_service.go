package services

import (
	"clean-architecture-api/lib"
	"clean-architecture-api/models"
	"clean-architecture-api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryService struct {
	logger          lib.Logger
	repository      repository.CategoryRepository
	paginationScope *gorm.DB
}

func NewCategoryService(logger lib.Logger, categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		logger:     logger,
		repository: categoryRepository,
	}
}

// PaginationScope
func (c CategoryService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) CategoryService {
	c.paginationScope = c.repository.WithTrx(c.repository.Scopes(scope)).DB
	return c
}
func (c CategoryService) GetAll() (map[string]interface{}, error) {
	var categories []models.Category
	var count int64

	err := c.repository.WithTrx(c.paginationScope).Find(&categories).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return gin.H{"data": categories, "count": count}, nil
}

func (c CategoryService) Create(category models.Category) error {
	return c.repository.Create(&category).Error
}
