package services

import (
	"clean-architecture-api/lib"
	"clean-architecture-api/models"
	"clean-architecture-api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BusinessService struct {
	repository      repository.BusinessRepository
	logger          lib.Logger
	paginationScope *gorm.DB
}

func NewBusinessService(repository repository.BusinessRepository, logger lib.Logger) *BusinessService {
	return &BusinessService{
		repository: repository,
		logger:     logger,
	}
}

// PaginationScope
func (c BusinessService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) BusinessService {
	c.paginationScope = c.repository.WithTrx(c.repository.Scopes(scope)).DB
	return c
}

func (bs BusinessService) GetAll() (map[string]interface{}, error) {
	var businesses []models.Business
	var count int64

	err := bs.repository.WithTrx(bs.paginationScope).Find(&businesses).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return gin.H{"data": businesses, "count": count}, nil
}

func (bs BusinessService) Create(business models.Business) error {
	result := bs.repository.Create(&business)
	if result.Error != nil {
		bs.logger.Error("Failed to save business")
		return result.Error
	}
	return nil
}
